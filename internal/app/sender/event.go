package sender

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/model"
)

// EventSender interface
type EventSender interface {
	Send(ctx context.Context, office *model.OfficeEvent) error
}
