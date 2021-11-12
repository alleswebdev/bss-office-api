package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const eventsTableName = "offices_events"

// EventRepo interface
type EventRepo interface {
	Add(ctx context.Context, event *model.OfficeEvent) error
	Lock(ctx context.Context, n uint64) ([]model.OfficeEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error
	Remove(ctx context.Context, eventIDs []uint64) error
}

type eventRepo struct {
	db *sqlx.DB
}

// NewEventRepo returns EventRepo interface
func NewEventRepo(db *sqlx.DB) EventRepo {
	return &eventRepo{db: db}
}

func (r *eventRepo) Add(ctx context.Context, event *model.OfficeEvent) error {
	query := database.StatementBuilder.
		Insert(eventsTableName).
		Columns("office_id", "type", "status", "payload", "created_at").
		Values(event.OfficeID, event.Type, event.Status, event.Payload, sq.Expr("NOW()")).
		Suffix("RETURNING id").
		RunWith(r.db)

	row := query.QueryRowContext(ctx)

	var id uint64

	err := row.Scan(&id)

	if err != nil {
		return errors.Wrap(err, "Add:Scan()")
	}

	event.ID = id

	return nil
}

func (r *eventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	sb := database.StatementBuilder.Delete(eventsTableName).Where(sq.Eq{"id": eventIDs})

	log.Info().Uints64("ids", eventIDs).Msg("ids")
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
		log.Debug().Uints64("ids", eventIDs).Msg("NO ROWS")
		return sql.ErrNoRows
	}

	return nil
}

func (r *eventRepo) Lock(ctx context.Context, n uint64) ([]model.OfficeEvent, error) {
	var events []model.OfficeEvent

	txErr := database.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) error {
		locked, err := database.AcquireTryLock(ctx, tx, database.LockTypeOfficeEvents, database.OfficeEventsTable)
		if err != nil {
			return errors.Wrap(err, "Lock()")
		}

		if !locked {
			return errors.Wrap(err, "not take lock")
		}

		sb := database.StatementBuilder.
			Update(eventsTableName).
			Prefix("with cte as (select id from offices_events where status <> ? order by id ASC limit ?)", model.Processed, n).
			Where(sq.Expr("exists (select * from cte where offices_events.id = cte.id)")).
			Set("status", model.Processed).
			Suffix("RETURNING id, office_id, type, status, created_at, payload")

		query, args, err := sb.ToSql()

		if err != nil {
			return errors.Wrap(err, "Lock: ToSql()")
		}

		err = r.db.SelectContext(ctx, &events, query, args...)

		if err != nil {
			return errors.Wrap(err, "Lock: SelectContext()")
		}

		return nil
	})

	return events, txErr
}

func (r *eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	sb := database.StatementBuilder.Update(eventsTableName).Where(sq.Eq{"id": eventIDs}).Set("Status", model.Deferred)

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
