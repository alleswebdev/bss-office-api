package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type lockType int
type entityID int

// LockTypeOfficeEvents - константа для блокировки таблицы office events
const (
	_ lockType = iota
	LockTypeOfficeEvents
)

// OfficeEventsTable - константа для блокировки таблицы office events
const (
	_ entityID = iota
	OfficeEventsTable
)

// AcquireTryLock берёт рекомендательную блокировку, которая снимается при завершении транзакции (xact)
func AcquireTryLock(ctx context.Context, tx *sqlx.Tx, lockID lockType, entityID entityID) (bool, error) {
	var isAcquired bool
	err := tx.GetContext(ctx, &isAcquired, fmt.Sprintf("select pg_try_advisory_xact_lock(%d, %d)", lockID, entityID))
	return isAcquired, err
}
