package model

import (
	"encoding/json"
	"errors"
	"time"
)

// EventType enum for event type
type EventType uint8

// EventStatus enum for event status
type EventStatus uint8

// Created - событие создано
// Updated - событие обновлено
// Removed - событие удалено
const (
	_ EventType = iota
	Created
	Updated
	Removed
)

// Deferred - событие заблокировано в репозитории для отправки
// Processed - событие обработанно
const (
	_ EventStatus = iota
	Deferred
	Processed
)

// OfficeEvent - office event model
type OfficeEvent struct {
	ID       uint64        `db:"id"`
	OfficeId uint64        `db:"office_id"`
	Type     EventType     `db:"type"`
	Status   EventStatus   `db:"status"`
	Created  time.Time     `db:"created"`
	Payload  OfficePayload `db:"payload"`
}

type OfficePayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Removed     bool   `json:"removed"`
}

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
