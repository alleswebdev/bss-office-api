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

const officesTableName = "offices"

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
		Select("id", "name", "description", "removed", "created", "updated").
		Where(sq.And{
			sq.Eq{"id": officeID},
			sq.NotEq{"removed": "true"},
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
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return &office, nil
}

// CreateOffice - create new office
func (r *repo) CreateOffice(ctx context.Context, office model.Office, tx *sqlx.Tx) (uint64, error) {

	sb := database.StatementBuilder.Insert(officesTableName).Columns(
		"name", "description").Values(office.Name, office.Description).Suffix("RETURNING id")

	query, args, err := sb.ToSql()

	if err != nil {
		return 0, errors.Wrap(err, "CreateOffice:ToSql()")
	}

	var queryer sqlx.QueryerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	rows, err := queryer.QueryContext(ctx, query, args...)

	if err != nil {
		return 0, errors.Wrap(err, "CreateOffice:QueryContext()")
	}

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, errors.Wrap(err, "CreateOffice:Scan()")
		}

		return id, nil
	}

	return 0, sql.ErrNoRows
}

//RemoveOffice - remove office by id
// office is not really delete, just set the removed flag to true
func (r *repo) RemoveOffice(ctx context.Context, officeID uint64, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).Set("removed", true).
		Where(sq.And{
			sq.Eq{"id": officeID},
			sq.NotEq{"removed": "true"},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "RemoveOffice:ToSql()")
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = r.db
	} else {
		execer = tx
	}

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
func (r *repo) UpdateOffice(ctx context.Context, officeID uint64, office model.Office, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set("name", office.Name).
		Set("description", office.Description).
		Where(sq.And{
			sq.Eq{"id": officeID},
			sq.NotEq{"removed": "true"},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOffice:ToSql()")
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = r.db
	} else {
		execer = tx
	}

	res, err := execer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, sql.ErrNoRows
	}

	return true, nil
}

//UpdateOfficeName - update  name field in office model by id
func (r *repo) UpdateOfficeName(ctx context.Context, officeID uint64, name string, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set("description", name).
		Where(sq.And{
			sq.Eq{"id": officeID},
			sq.NotEq{"removed": "true"},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOfficeName:ToSql()")
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = r.db
	} else {
		execer = tx
	}

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

//UpdateOfficeDescription - updated all description field in office model by id
func (r *repo) UpdateOfficeDescription(ctx context.Context, officeID uint64, description string, tx *sqlx.Tx) (bool, error) {
	sb := database.StatementBuilder.
		Update(officesTableName).
		Set("description", description).
		Where(sq.And{
			sq.Eq{"id": officeID},
			sq.NotEq{"removed": "true"},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOfficeDescription:ToSql()")
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = r.db
	} else {
		execer = tx
	}

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

// ListOffices - return all offices
func (r *repo) ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error) {
	sb := database.StatementBuilder.
		Select("id", "name", "description", "removed", "created", "updated").
		From(officesTableName).
		Where(sq.NotEq{"removed": "true"}).
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
