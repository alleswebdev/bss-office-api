package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/model"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
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

const (
	officesEventsIdColumn        = "id"
	officesEventsOfficeIdColumn  = "office_id"
	officesEventsTypeColumn      = "type"
	officesEventsStatusColumn    = "status"
	officesEventsPayloadColumn   = "payload"
	officesEventsCreatedAtColumn = "created_at"
)

// NewEventRepo returns EventRepo interface
func NewEventRepo(db *sqlx.DB) EventRepo {
	return &eventRepo{db: db}
}

func (r *eventRepo) Add(ctx context.Context, event *model.OfficeEvent) error {
	payload, err := convertBssOfficeToJsonb(&event.Payload)

	if err != nil {
		return errors.Wrap(err, "Add()")
	}

	query := database.StatementBuilder.
		Insert(eventsTableName).
		Columns(
			officesEventsIdColumn,
			officesEventsTypeColumn,
			officesEventsStatusColumn,
			officesEventsPayloadColumn,
			officesEventsCreatedAtColumn).
		Values(event.OfficeID, event.Type, event.Status, payload, sq.Expr("NOW()")).
		Suffix("RETURNING " + officesEventsIdColumn).
		RunWith(r.db)

	row := query.QueryRowContext(ctx)

	var id uint64

	err = row.Scan(&id)

	if err != nil {
		return errors.Wrap(err, "Add:Scan()")
	}

	event.ID = id

	return nil
}

func (r *eventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	sb := database.StatementBuilder.Delete(eventsTableName).Where(sq.Eq{officesEventsIdColumn: eventIDs})

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
		return ErrOfficeNotFound
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
			Prefix("WITH cte as (SELECT id FROM offices_events WHERE status <> ? ORDER BY id ASC LIMIT ?)", model.Processed, n).
			Where(sq.Expr("EXISTS (SELECT * FROM cte WHERE offices_events.id = cte.id)")).
			Set("status", model.Processed).
			Suffix("RETURNING *")

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
	sb := database.StatementBuilder.Update(eventsTableName).
		Where(sq.Eq{officesEventsIdColumn: eventIDs}).
		Set(officesEventsStatusColumn, model.Deferred)

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

func convertBssOfficeToJsonb(o *model.OfficePayload) ([]byte, error) {
	var pbStream = &pb.Office{
		Id:          o.ID,
		Name:        o.Name,
		Description: o.Description,
		Removed:     o.Removed,
	}

	payload, err := protojson.Marshal(pbStream)

	if err != nil {
		return nil, errors.Wrap(err, "convertBssOfficeToJsonb()")
	}

	return payload, nil
}
