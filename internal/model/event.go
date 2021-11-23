package model

import (
	"database/sql"
	"encoding/json"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// EventType enum for event type
//go:generate stringer -linecomment -type=EventType
type EventType uint8

// EventStatus enum for event status
type EventStatus uint8

// Created - событие создано
// Updated - событие обновлено
// Removed - событие удалено
const (
	_                        EventType = iota
	Created                            //created
	Updated                            //updated
	Removed                            //removed
	OfficeNameUpdated                  //office_name_updated
	OfficeDescriptionUpdated           //office_description_updated
)

// Deferred - событие ожидает обработки
// Processed - событие обрабатывается
const (
	_ EventStatus = iota
	Deferred
	Processed
)

// OfficeEvent - office event model
type OfficeEvent struct {
	ID       uint64        `db:"id"`
	OfficeID uint64        `db:"office_id"`
	Type     EventType     `db:"type"`
	Status   EventStatus   `db:"status"`
	Created  time.Time     `db:"created_at"`
	Updated  sql.NullTime  `db:"updated_at"`
	Payload  OfficePayload `db:"payload"`
}

//OfficePayload Сктура для записи информации о изменениях в сущности office
type OfficePayload struct {
	ID          uint64 `json:"id,string"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Removed     bool   `json:"removed,omitempty"`
}

// Scan - кастомный сканер для OfficePayload
func (op *OfficePayload) Scan(src interface{}) (err error) {
	var payload OfficePayload
	if src == nil {
		return nil
	}

	switch src.(type) {
	case string:
		err = json.Unmarshal([]byte(src.(string)), &payload)
	case []byte:
		err = json.Unmarshal(src.([]byte), &payload)
	default:
		return errors.New("incompatible type")
	}

	if err != nil {
		return err
	}

	*op = payload

	return nil
}

func ConvertBssOfficeEventToPb(o *OfficeEvent) *pb.OfficeEvent {
	pb := &pb.OfficeEvent{
		Id:       o.ID,
		OfficeId: o.OfficeID,
		Status:   uint64(o.Status),
		Type:     o.Type.String(),
		Created:  timestamppb.New(o.Created),
		Updated:  timestamppb.New(o.Updated.Time),
		Payload:  ConvertBssOfficePayloadToPb(&o.Payload),
	}

	return pb
}

func ConvertBssOfficePayloadToPb(op *OfficePayload) *pb.OfficePayload {
	pb := &pb.OfficePayload{
		Id:          op.ID,
		Name:        op.Name,
		Description: op.Description,
	}

	return pb
}
