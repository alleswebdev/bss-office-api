package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/pkg/errors"
)

const tableName = "offices_events"

// EventRepo interface
type EventRepo interface {
	Lock(ctx context.Context, n uint64) ([]model.OfficeEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error
	Remove(ctx context.Context, eventIDs []uint64) error
}

type eventRepo struct {
	db *sqlx.DB
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB) EventRepo {
	return &eventRepo{db: db}
}

// StatementBuilder глобальная переменная с сконфигурированным плейсхолдером для pgsql
var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (r *eventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	sb := StatementBuilder.Delete(tableName).Where(sq.Eq{"id": eventIDs})

	query, args, err := sb.ToSql()

	if err != nil {
		return errors.Wrap(err, "Remove: ToSql()")
	}

	res, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.Wrap(err, "Remove: ExecContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return errors.Wrap(err, "Remove: RowsAffected()")
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *eventRepo) Lock(ctx context.Context, n uint64) ([]model.OfficeEvent, error) {
	sb := StatementBuilder.
		Update(tableName).
		Prefix("with cte as (select id from offices_events where status <> ? limit ? order by id ASC)", model.Processed, n).
		Where(sq.Expr("exists (select * from cte where offices_events.id = cte.id)")).
		Set("status", model.Processed).
		Suffix("RETURNING id, office_id, type,status,created, payload")

	query, args, err := sb.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "Lock: ToSql()")
	}

	var events []model.OfficeEvent
	err = r.db.SelectContext(ctx, &events, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "Lock: SelectContext()")
	}

	return events, nil
}

func (r *eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	sb := StatementBuilder.Update(tableName).Where(sq.Eq{"id": eventIDs}).Set("Status", model.Deferred)

	query, args, err := sb.ToSql()

	if err != nil {
		return errors.Wrap(err, "Unlock: ToSql()")
	}

	res, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.Wrap(err, "Unlock: ExecContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return errors.Wrap(err, "Unlock: RowsAffected()")
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}
