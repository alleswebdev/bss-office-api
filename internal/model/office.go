package model

import "time"

type Office struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Created     time.Time `db:"created"`
}
