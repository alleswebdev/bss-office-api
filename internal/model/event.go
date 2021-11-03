package model

// EventType event type on of:
type EventType uint8

// EventStatus enum for event status
type EventStatus uint8

// Created - событие создано
// Updated - событие обновлено
// Removed - событие удалено
// Deferred - событие заблокировано в репозитории для отправки
// Processed - событие обработанно
const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

// OfficeEvent - office event model
type OfficeEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Office
}
