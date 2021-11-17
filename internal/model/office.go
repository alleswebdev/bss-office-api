package model

import (
	"database/sql"
	"time"
)

// Office model for office
type Office struct {
	ID          uint64       `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	Removed     bool         `db:"removed"`
	Updated     sql.NullTime `db:"updated_at"`
	Created     time.Time    `db:"created_at"`
}
