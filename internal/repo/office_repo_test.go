package repo

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/require"
	"regexp"
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

	db, mock, err := sqlmock.New()
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

	f.dbMock.ExpectQuery(`INSERT INTO offices (.+) VALUES (.+) RETURNING id`).
		WithArgs(testOffice.Name, testOffice.Description).
		WillReturnRows(rows)

	resultID, err := f.officeRepo.CreateOffice(context.Background(), testOffice, nil)

	require.NoError(t, err)
	require.Equal(t, testOffice.ID, resultID)
}

func Test_repo_DescribeOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "removed", "created_at", "updated_at"}).
		AddRow(1, testOffice.Name, testOffice.Description, false, time.Now(), time.Now())

	f.dbMock.ExpectQuery(`SELECT (.+) FROM offices WHERE \(id = \$1 AND removed <> \$2\) LIMIT 1`).
		WithArgs(testOffice.ID, true).
		WillReturnRows(rows)

	result, err := f.officeRepo.DescribeOffice(context.Background(), testOffice.ID)

	require.NoError(t, err)
	require.Equal(t, result.ID, testOffice.ID)
	require.Equal(t, result.Name, testOffice.Name)
	require.Equal(t, result.Description, testOffice.Description)
}

func Test_repo_DescribeOffice_Err_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	f.dbMock.ExpectQuery(`SELECT (.+) FROM offices WHERE \(id = \$1 AND removed <> \$2\) LIMIT 1`).
		WithArgs(testOffice.ID, true).
		WillReturnError(sql.ErrNoRows)

	_, err := f.officeRepo.DescribeOffice(context.Background(), testOffice.ID)

	require.ErrorIs(t, err, ErrOfficeNotFound)
}

func Test_repo_ListOffices(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "removed", "created_at", "updated_at"}).
		AddRow(1, testOffice.Name, testOffice.Description, false, time.Now(), time.Now())

	f.dbMock.ExpectQuery(`SELECT (.+) FROM offices WHERE removed <> \$1 LIMIT 0 OFFSET 5`).
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

	expectSql := regexp.QuoteMeta(`UPDATE offices SET removed = $1 WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(true, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.RemoveOffice(context.Background(), testOffice.ID, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_RemoveOffice_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET removed = $1 WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(true, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(0, 0))

	result, err := f.officeRepo.RemoveOffice(context.Background(), testOffice.ID, nil)

	require.NoError(t, err)
	require.Equal(t, result, false)
}

func Test_repo_UpdateOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET name = $1, description = $2 WHERE (id = $3 AND removed <> $4)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(testOffice.Name, testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOffice(context.Background(), testOffice.ID, testOffice, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOffice_Err_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET name = $1, description = $2 WHERE (id = $3 AND removed <> $4)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(testOffice.Name, testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(0, 0))

	result, err := f.officeRepo.UpdateOffice(context.Background(), testOffice.ID, testOffice, nil)

	require.ErrorIs(t, err, ErrOfficeNotFound)
	require.Equal(t, result, false)
}

func Test_repo_UpdateOfficeDescription(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET description = $1 WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOfficeDescription(context.Background(), testOffice.ID, testOffice.Description, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOfficeDescription_Err_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET description = $1 WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(0, 0))

	result, err := f.officeRepo.UpdateOfficeDescription(context.Background(), testOffice.ID, testOffice.Description, nil)

	require.ErrorIs(t, err, ErrOfficeNotFound)
	require.Equal(t, result, false)
}

func Test_repo_UpdateOfficeName(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET name = $1 WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOfficeName(context.Background(), testOffice.ID, testOffice.Name, nil)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOfficeName_Err_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSql := regexp.QuoteMeta(`UPDATE offices SET name = $1 WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSql).
		WithArgs(testOffice.Name, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(0, 0))

	result, err := f.officeRepo.UpdateOfficeName(context.Background(), testOffice.ID, testOffice.Name, nil)

	require.ErrorIs(t, err, ErrOfficeNotFound)
	require.Equal(t, result, false)
}
