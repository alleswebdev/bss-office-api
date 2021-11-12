package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

type eventRepoFixture struct {
	eventRepo EventRepo
	db        *sqlx.DB
	dbMock    sqlmock.Sqlmock
}

var testEventModel = model.OfficeEvent{
	OfficeID: 1,
	Type:     model.Created,
	Status:   model.Deferred,
	Created:  time.Now(),
}

func setUpEventRepo(t *testing.T) eventRepoFixture {
	var fixture eventRepoFixture

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	fixture.db = sqlx.NewDb(db, "sqlmock")
	fixture.dbMock = mock
	fixture.eventRepo = NewEventRepo(fixture.db)

	return fixture
}

func (f *eventRepoFixture) tearDown() {
	f.db.Close()
}

func Test_eventRepo_Lock(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id", "office_id", "type", "status", "created_at", "payload"}).
		AddRow(1, 1, model.Created, model.Processed, time.Now(), "{}").
		AddRow(2, 2, model.Updated, model.Processed, time.Now(), "{}").
		AddRow(3, 3, model.Removed, model.Processed, time.Now(), "{}")

	testLimit := uint64(3)

	f.dbMock.ExpectBegin()
	lockSql := regexp.QuoteMeta("select pg_try_advisory_xact_lock(1, 1)")

	f.dbMock.ExpectQuery(lockSql).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(1))

	expectSql := regexp.QuoteMeta(`
WITH cte as
	(SELECT id
	FROM offices_events 
	WHERE status <> $1 
	ORDER BY id 
	ASC LIMIT $2)
UPDATE offices_events SET status = $3 
WHERE EXISTS (SELECT * FROM cte WHERE offices_events.id = cte.id) 
RETURNING id, office_id, type, status, created_at, payload`)

	f.dbMock.ExpectQuery(expectSql).
		WithArgs(model.Processed, testLimit, model.Processed).
		WillReturnRows(rows)

	f.dbMock.ExpectCommit()

	result, err := f.eventRepo.Lock(context.Background(), testLimit)

	require.NoError(t, err)
	require.Equal(t, testLimit, uint64(len(result)))
}

func Test_eventRepo_Remove(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`DELETE FROM offices_events WHERE id IN (.+)`).
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := f.eventRepo.Remove(context.Background(), []uint64{1, 2})

	require.NoError(t, err)
}

func Test_eventRepo_Unlock(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices_events SET Status = \$1 WHERE id IN (.+)`).
		WithArgs(model.Deferred, 1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := f.eventRepo.Unlock(context.Background(), []uint64{1, 2})

	require.NoError(t, err)
}

func Test_eventRepo_Add(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	expectPayload, err := convertBssOfficeToJsonb(&testEventModel.Payload)

	require.NoError(t, err)

	f.dbMock.ExpectQuery(`INSERT INTO offices_events  (.+) VALUES  (.+) RETURNING id`).
		WithArgs(testEventModel.OfficeID, testEventModel.Type, testEventModel.Status, expectPayload).
		WillReturnRows(rows)

	err = f.eventRepo.Add(context.Background(), &testEventModel)

	require.NoError(t, err)
}
