package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/ozonmp/bss-office-api/internal/model"
)

// Repo is DAO for Office
type Repo interface {
	DescribeOffice(ctx context.Context, templateID uint64) (*model.Office, error)
	CreateOffice(ctx context.Context, office model.Office) (uint64, error)
	RemoveOffice(ctx context.Context, officeID uint64) (bool, error)
	ListOffices(ctx context.Context) ([]*model.Office, error)
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
	log.Debug().Uint64("officeID", officeID).Msg("DescribeOffice")

	return nil, nil
}

// CreateOffice - create new office
func (r *repo) CreateOffice(ctx context.Context, office model.Office) (uint64, error) {
	log.Debug().Str("name", office.Name).Str("description", office.Description).Msg("CreateOffice")

	return 0, nil
}

//RemoveOffice - remove office by id
func (r *repo) RemoveOffice(ctx context.Context, officeID uint64) (bool, error) {
	log.Debug().Uint64("officeID", officeID).Msg("RemoveOffice")

	return true, nil
}

// ListOffices - return all offices
func (r *repo) ListOffices(ctx context.Context) ([]*model.Office, error) {
	log.Debug().Msg("ListOffices")

	return []*model.Office{}, nil
}
