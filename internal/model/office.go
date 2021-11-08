package model

import "time"

// Office model for office
type Office struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Removed     bool      `db:"removed"`
	Updated     time.Time `db:"updated"`
	Created     time.Time `db:"created"`
}
