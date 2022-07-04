package sqlx_repo

import (
	"context"

	"github.com/chillyNick/librarySearch/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type repository struct {
	db *sqlx.DB
}

// New returns sqlx based repository
func New(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetBooks(ctx context.Context, author string) (books []models.Book, err error) {
	const query = `
		SELECT b.* FROM author a
			JOIN author_book ab ON a.id = ab.author_id
			JOIN book b ON b.id = ab.book_id
		WHERE a.name = ?
	`

	err = r.db.SelectContext(ctx, &books, query, author)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get books by author name: %s", author)

		return
	}

	return
}

func (r *repository) GetAuthors(ctx context.Context, book string) (authors []models.Author, err error) {
	const query = `
		SELECT a.* FROM book b
			JOIN author_book ab ON b.id = ab.book_id
			JOIN author a ON a.id = ab.author_id
		WHERE b.title = ?
	`

	err = r.db.SelectContext(ctx, &authors, query, book)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get authors by book title: %s", book)

		return
	}

	return
}
