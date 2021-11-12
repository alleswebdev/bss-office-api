package sender

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/rs/zerolog/log"
)

// EventSender interface
type EventSender interface {
	Send(ctx context.Context, office *model.OfficeEvent) error
}

type dummySender struct {
}

func NewDummySender() *dummySender {
	return &dummySender{}
}

func (s *dummySender) Send(_ context.Context, office *model.OfficeEvent) error {
	log.Debug().Uint64("ID", office.ID).Msg("sending event...")

	return nil
}
