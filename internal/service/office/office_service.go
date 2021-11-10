package office

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/repo"
	"github.com/pkg/errors"
)

// EventRepo - часть публичного интерфейса, которая используется в этом сервисе
type EventRepo interface {
	Add(ctx context.Context, event *model.OfficeEvent) error
	Lock(ctx context.Context, n uint64) ([]model.OfficeEvent, error)
}

type officeService struct {
	officeRepo repo.OfficeRepo
	eventRepo  EventRepo
	db         *sqlx.DB
}

// NewOfficeService создаёт новый сервис для работы с моделью office и создания событий о её изменениях
func NewOfficeService(or repo.OfficeRepo, er EventRepo, db *sqlx.DB) *officeService {
	return &officeService{
		officeRepo: or,
		eventRepo:  er,
		db:         db,
	}
}

func (s *officeService) DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error) {
	return s.officeRepo.DescribeOffice(ctx, officeID)
}

func (s *officeService) ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error) {
	return s.officeRepo.ListOffices(ctx, limit, offset)
}

func (s *officeService) RemoveOffice(ctx context.Context, officeID uint64) (bool, error) {
	var result = false

	txErr := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		result, err = s.officeRepo.RemoveOffice(ctx, officeID, tx)

		if err != nil {
			return err
		}

		if !result {
			return err
		}

		err = s.eventRepo.Add(ctx, &model.OfficeEvent{
			OfficeID: officeID,
			Type:     model.Removed,
			Status:   model.Deferred,
			Payload: model.OfficePayload{
				ID:      officeID,
				Removed: true,
			},
		})

		if err != nil {
			return errors.Wrap(err, "RemoveOffice() : cannot add event")
		}

		return nil
	})

	return result, txErr
}

func (s *officeService) CreateOffice(ctx context.Context, office model.Office) (uint64, error) {
	var officeID = uint64(0)

	txErr := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		officeID, err = s.officeRepo.CreateOffice(ctx, office, tx)

		if err != nil {
			return err
		}

		err = s.eventRepo.Add(ctx, &model.OfficeEvent{
			OfficeID: officeID,
			Type:     model.Created,
			Status:   model.Deferred,
			Payload: model.OfficePayload{
				ID:          officeID,
				Name:        office.Name,
				Description: office.Description,
				Removed:     office.Removed,
			},
		})

		if err != nil {
			return errors.Wrap(err, "CreateOffice() : cannot add event")
		}

		return nil
	})

	return officeID, txErr
}

func (s *officeService) UpdateOffice(ctx context.Context, officeID uint64, office model.Office) (bool, error) {
	var result = false

	txErr := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		result, err = s.officeRepo.UpdateOffice(ctx, officeID, office, tx)

		if err != nil {
			return err
		}

		if !result {
			return err
		}

		err = s.eventRepo.Add(ctx, &model.OfficeEvent{
			OfficeID: officeID,
			Type:     model.Updated,
			Status:   model.Deferred,
			Payload: model.OfficePayload{
				ID:          officeID,
				Name:        office.Name,
				Description: office.Description,
			},
		})

		if err != nil {
			return errors.Wrap(err, "UpdateOffice() : cannot add event")
		}

		return nil
	})

	return result, txErr
}

func (s *officeService) UpdateOfficeName(ctx context.Context, officeID uint64, name string) (bool, error) {
	var result = false

	txErr := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		result, err = s.officeRepo.UpdateOfficeName(ctx, officeID, name, tx)

		if err != nil {
			return err
		}

		if !result {
			return err
		}

		err = s.eventRepo.Add(ctx, &model.OfficeEvent{
			OfficeID: officeID,
			Type:     model.OfficeNameUpdated,
			Status:   model.Deferred,
			Payload: model.OfficePayload{
				ID:   officeID,
				Name: name,
			},
		})

		if err != nil {
			return errors.Wrap(err, "UpdateOfficeName() : cannot add event")
		}

		return nil
	})

	return result, txErr
}

func (s *officeService) UpdateOfficeDescription(ctx context.Context, officeID uint64, description string) (bool, error) {
	var result = false

	txErr := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		result, err = s.officeRepo.UpdateOfficeName(ctx, officeID, description, tx)

		if err != nil {
			return err
		}

		if !result {
			return err
		}

		err = s.eventRepo.Add(ctx, &model.OfficeEvent{
			OfficeID: officeID,
			Type:     model.OfficeDescriptionUpdated,
			Status:   model.Deferred,
			Payload: model.OfficePayload{
				ID:          officeID,
				Description: description,
			},
		})

		if err != nil {
			return errors.Wrap(err, "UpdateOfficeDescription() : cannot add event")
		}

		return nil
	})

	return result, txErr
}
