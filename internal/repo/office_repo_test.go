package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var testOffice = model.Office{
	ID:          uint64(1),
	Name:        "test",
	Description: "test",
}

type officeRepoFixture struct {
	officeRepo OfficeRepo
	db         *sqlx.DB
	dbMock     sqlmock.Sqlmock
}

func setUp(t *testing.T) officeRepoFixture {
	var fixture officeRepoFixture

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	fixture.db = sqlx.NewDb(db, "sqlmock")
	fixture.dbMock = mock
	fixture.officeRepo = NewOfficeRepo(fixture.db)

	return fixture
}

func (f *officeRepoFixture) tearDown() {
	f.db.Close()
}

func Test_repo_CreateOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	f.dbMock.ExpectQuery(`INSERT INTO offices (name,description) VALUES ($1,$2) RETURNING id`).
		WithArgs(testOffice.Name, testOffice.Description).
		WillReturnRows(rows)

	resultID, err := f.officeRepo.CreateOffice(context.Background(), testOffice, nil)

	require.NoError(t, err)
	require.Equal(t, testOffice.ID, resultID)
}

func Test_repo_DescribeOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "removed", "created", "updated"}).
		AddRow(1, testOffice.Name, testOffice.Description, false, time.Now(), time.Now())

	f.dbMock.ExpectQuery(`SELECT id, name, description, removed, created, updated FROM offices WHERE (id = $1 AND removed <> $2) LIMIT 1`).
		WithArgs(testOffice.ID, true).
		WillReturnRows(rows)

	result, err := f.officeRepo.DescribeOffice(context.Background(), testOffice.ID)

	require.NoError(t, err)
	require.Equal(t, result.ID, testOffice.ID)
	require.Equal(t, result.Name, testOffice.Name)
	require.Equal(t, result.Description, testOffice.Description)
}

func Test_repo_ListOffices(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "removed", "created", "updated"}).
		AddRow(1, testOffice.Name, testOffice.Description, false, time.Now(), time.Now())

	f.dbMock.ExpectQuery(`SELECT id, name, description, removed, created, updated FROM offices WHERE removed <> $1 LIMIT 0 OFFSET 5`).
		WithArgs(true).
		WillReturnRows(rows)

	res, err := f.officeRepo.ListOffices(context.Background(), 0, 5)

	require.Equal(t, testOffice.Name, res[0].Name)
	require.Equal(t, testOffice.Description, res[0].Description)
	require.NoError(t, err)
}

func Test_repo_RemoveOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices SET removed = $1 WHERE (id = $2 AND removed <> $3)`).
		WithArgs(true, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.RemoveOffice(context.Background(), testOffice.ID, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices SET name = $1, description = $2 WHERE (id = $3 AND removed <> $4)`).
		WithArgs(testOffice.Name, testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOffice(context.Background(), testOffice.ID, testOffice, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOfficeDescription(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices SET description = $1 WHERE (id = $2 AND removed <> $3)`).
		WithArgs(testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOfficeDescription(context.Background(), testOffice.ID, testOffice.Description, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOfficeName(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices SET name = $1 WHERE (id = $2 AND removed <> $3)`).
		WithArgs(testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOfficeName(context.Background(), testOffice.ID, testOffice.Name, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}
