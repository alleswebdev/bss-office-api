package repo

import (
	"github.com/ozonmp/bss-office-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.OfficeEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.OfficeEvent) error
	Remove(eventIDs []uint64) error
}