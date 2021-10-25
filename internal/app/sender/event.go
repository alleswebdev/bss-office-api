package sender

import (
	"github.com/ozonmp/bss-office-api/internal/model"
)

type EventSender interface {
	Send(office *model.OfficeEvent) error
}
