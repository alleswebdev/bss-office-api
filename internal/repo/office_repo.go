package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/pkg/errors"
)

var ErrOfficeNotFound = errors.New("office not found")

const officesTableName = "offices"

const (
	officeIdColumn          = "id"
	officeNameColumn        = "name"
	officeDescriptionColumn = "description"
	officeRemovedColumn     = "removed"
	officeCreatedAtColumn   = "created_at"
	officeUpdatedAtColumn   = "updated_at"
)

// OfficeRepo is DAO for Office
type OfficeRepo interface {
	DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error)
	CreateOffice(ctx context.Context, office model.Office, tx *sqlx.Tx) (uint64, error)
	RemoveOffice(ctx context.Context, officeID uint64, tx *sqlx.Tx) (bool, error)
	ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error)
	UpdateOffice(ctx context.Context, officeID uint64, office model.Office, tx *sqlx.Tx) (bool, error)
	UpdateOfficeName(ctx context.Context, officeID uint64, name string, tx *sqlx.Tx) (bool, error)
	UpdateOfficeDescription(ctx context.Context, officeID uint64, description string, tx *sqlx.Tx) (bool, error)
}

type repo struct {
	db *sqlx.DB
}

// NewOfficeRepo returns OfficeRepo interface
func NewOfficeRepo(db *sqlx.DB) OfficeRepo {
	return &repo{db: db}
}

// DescribeOffice Describe an office by id
func (r *repo) DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error) {
	sb := database.StatementBuilder.
		Select(
			officeIdColumn,
			officeNameColumn,
			officeDescriptionColumn,
			officeRemovedColumn,
			officeCreatedAtColumn,
			officeUpdatedAtColumn).
		Where(sq.And{
			sq.Eq{officeIdColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		}).
		From(officesTableName).
		Limit(1)

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "DescribeOffice:ToSql()")
	}

	var office model.Office
	err = r.db.GetContext(ctx, &office, query, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrOfficeNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return &office, nil
}

// CreateOffice - create new office
func (r *repo) CreateOffice(ctx context.Context, office model.Office, tx *sqlx.Tx) (uint64, error) {

	sb := database.StatementBuilder.
		Insert(officesTableName).
		Columns(officeNameColumn, officeDescriptionColumn).
		Values(office.Name, office.Description).
		Suffix("RETURNING " + officeIdColumn)

	query, args, err := sb.ToSql()

	if err != nil {
		return 0, errors.Wrap(err, "CreateOffice:ToSql()")
	}

	queryer := r.getQueryer(tx)
	row := queryer.QueryRowxContext(ctx, query, args...)

	var id uint64
	err = row.Scan(&id)

	if err != nil {
		return 0, errors.Wrap(err, "CreateOffice:Scan()")
	}

	return id, nil
}

//RemoveOffice - remove office by id
// office is not really delete, just set the removed flag to true
func (r *repo) RemoveOffice(ctx context.Context, officeID uint64, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set(officeRemovedColumn, true).
		Where(sq.And{
			sq.Eq{officeIdColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "RemoveOffice:ToSql()")
	}

	execer := r.getExecer(tx)
	res, err := execer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, nil
	}

	return true, nil
}

//UpdateOffice - update all editable fields in office model by id
func (r *repo) UpdateOffice(ctx context.Context, officeID uint64, officeModel model.Office, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set(officeNameColumn, officeModel.Name).
		Set(officeDescriptionColumn, officeModel.Description).
		Where(sq.And{
			sq.Eq{officeIdColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOffice:ToSql()")
	}

	execer := r.getExecer(tx)
	res, err := execer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, ErrOfficeNotFound
	}

	return true, nil
}

//UpdateOfficeName - update  name field in office model by id
func (r *repo) UpdateOfficeName(ctx context.Context, officeID uint64, name string, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set(officeNameColumn, name).
		Where(sq.And{
			sq.Eq{officeIdColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOfficeName:ToSql()")
	}

	execer := r.getExecer(tx)
	res, err := execer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, ErrOfficeNotFound
	}

	return true, nil
}

//UpdateOfficeDescription - updated all description field in office model by id
func (r *repo) UpdateOfficeDescription(ctx context.Context, officeID uint64, description string, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set(officeDescriptionColumn, description).
		Where(sq.And{
			sq.Eq{officeIdColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOfficeDescription:ToSql()")
	}

	execer := r.getExecer(tx)
	res, err := execer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, ErrOfficeNotFound
	}

	return true, nil
}

// ListOffices - return all offices
func (r *repo) ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error) {
	sb := database.StatementBuilder.
		Select(
			officeIdColumn,
			officeNameColumn,
			officeDescriptionColumn,
			officeRemovedColumn,
			officeCreatedAtColumn,
			officeUpdatedAtColumn).
		From(officesTableName).
		OrderBy(officeIdColumn).
		Where(sq.NotEq{officeRemovedColumn: true}).
		Limit(limit).Offset(offset)

	query, args, err := sb.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "ListOffices:ToSql()")
	}

	var offices []*model.Office

	err = r.db.SelectContext(ctx, &offices, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return offices, nil
}

func (r *repo) getQueryer(tx *sqlx.Tx) sqlx.QueryerContext {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *repo) getExecer(tx *sqlx.Tx) sqlx.ExecerContext {
	if tx == nil {
		return r.db
	}
	return tx
}
