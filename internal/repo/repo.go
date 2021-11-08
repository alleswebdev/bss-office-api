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

const tableName = "offices"

// Repo is DAO for Office
type Repo interface {
	DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error)
	CreateOffice(ctx context.Context, office model.Office) (uint64, error)
	RemoveOffice(ctx context.Context, officeID uint64) (bool, error)
	ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

// DescribeOffice Describe an office by id
func (r *repo) DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error) {
	sb := database.StatementBuilder.
		Select("id", "name", "description", "removed", "created", "updated").
		Where(sq.Eq{"id": officeID}).
		From(tableName).
		Limit(1)

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
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
func (r *repo) CreateOffice(ctx context.Context, office model.Office) (uint64, error) {
	query := database.StatementBuilder.Insert(tableName).Columns(
		"name", "description").Values(office.Name, office.Description).Suffix("RETURNING id").RunWith(r.db)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return 0, err
	}

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	}

	return 0, sql.ErrNoRows
}

//RemoveOffice - remove office by id
// office is not really delete, just set the removed flag to true
func (r *repo) RemoveOffice(ctx context.Context, officeID uint64) (bool, error) {
	sb := database.StatementBuilder.
		Update(tableName).Set("removed", true).
		Where(sq.And{
			sq.Eq{"id": officeID},
			sq.NotEq{"removed": "true"},
		}).RunWith(r.db)

	query, args, err := sb.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.db.ExecContext(ctx, query, args...)

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
		From(tableName).
		Where(sq.NotEq{"removed": "true"}).
		Limit(limit).Offset(offset)

	sql, args, err := sb.ToSql()

	if err != nil {
		return nil, err
	}

	var offices []*model.Office

	err = r.db.SelectContext(ctx, &offices, sql, args...)

	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return offices, nil
}
