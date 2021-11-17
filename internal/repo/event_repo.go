package repo

import (
	"context"
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

// ErrNoneRowsUnlock ошибка, возникающая когда при разблокировании событий
//не было изменено ни одного запрошенного особытия
var ErrNoneRowsUnlock = errors.New("no have affected rows on unlock")

type eventRepo struct {
	db *sqlx.DB
}

const (
	officesEventsIDColumn        = "id"
	officesEventsOfficeIDColumn  = "office_id"
	officesEventsTypeColumn      = "type"
	officesEventsStatusColumn    = "status"
	officesEventsPayloadColumn   = "payload"
	officesEventsCreatedAtColumn = "created_at"
	officesEventsUpdatedAtColumn = "updated_at"
)

// NewEventRepo returns EventRepo interface
func NewEventRepo(db *sqlx.DB) EventRepo {
	return &eventRepo{db: db}
}

func (r *eventRepo) Add(ctx context.Context, event *model.OfficeEvent) error {
	payload, err := convertBssOfficeToJsonb(&event.Payload)

	if err != nil {
		return errors.Wrap(err, "convertBssOfficeToJsonb()")
	}

	query := database.StatementBuilder.
		Insert(eventsTableName).
		Columns(
			officesEventsOfficeIDColumn,
			officesEventsTypeColumn,
			officesEventsStatusColumn,
			officesEventsPayloadColumn,
			officesEventsCreatedAtColumn).
		Values(event.OfficeID, event.Type, event.Status, payload, sq.Expr("NOW()")).
		Suffix("RETURNING " + officesEventsIDColumn).
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
	sb := database.StatementBuilder.Delete(eventsTableName).Where(sq.Eq{officesEventsIDColumn: eventIDs})

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

func (r *eventRepo) Lock(ctx context.Context, batchSize uint64) ([]model.OfficeEvent, error) {
	whereSubquery := database.StatementBuilder.
		Select(officesEventsIDColumn).
		From(eventsTableName).
		Where(sq.Eq{officesEventsStatusColumn: model.Deferred}).
		OrderBy(officesEventsIDColumn).
		Limit(batchSize).
		Suffix("FOR UPDATE SKIP LOCKED")

	sb := database.StatementBuilder.
		Update(eventsTableName).
		Where(whereSubquery.Prefix(officesEventsIDColumn+" IN (").Suffix(")")).
		Set(officesEventsStatusColumn, model.Processed).
		Set(officesEventsUpdatedAtColumn, sq.Expr("NOW()")).
		Suffix("RETURNING *")

	query, args, err := sb.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "Lock: ToSql()")
	}

	events := make([]model.OfficeEvent, 0, batchSize)
	err = r.db.SelectContext(ctx, &events, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "Lock: SelectContext()")
	}

	return events, nil
}

func (r *eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	sb := database.StatementBuilder.Update(eventsTableName).
		Where(sq.Eq{officesEventsIDColumn: eventIDs}).
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
		return ErrNoneRowsUnlock
	}

	return nil
}

func convertBssOfficeToJsonb(o *model.OfficePayload) ([]byte, error) {
	pbStream := &pb.Office{
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
