package model

// Office model for office
type Office struct {
	ID          uint64 `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
