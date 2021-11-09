package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/pkg/errors"
)

const eventsTableName = "offices_events"

// EventRepo interface
type EventRepo interface {
	Add(ctx context.Context, event *model.OfficeEvent) error
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
		Columns("office_id", "type", "status", "created").
		Values(event.OfficeId, event.Type, event.Status, sq.Expr("NOW()")).
		Suffix("RETURNING id").
		RunWith(r.db)

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return errors.Wrap(err, "Add:QueryContext()")
	}

	var id uint64

	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return errors.Wrap(err, "Add:Scan()")
		}
	}

	event.ID = id

	return nil
}
