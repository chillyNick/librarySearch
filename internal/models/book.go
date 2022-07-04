package models

// Book - book entity
type Book struct {
	ID    int32  `db:"id"`
	Title string `db:"title"`
}
