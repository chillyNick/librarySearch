package models

// Author - author entity
type Author struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}
