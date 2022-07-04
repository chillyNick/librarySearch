package repo

import (
	"context"

	"github.com/chillyNick/librarySearch/internal/models"
)

//Repo is interface for working with library db
type Repo interface {
	//GetBooks returns books entities by author name from db
	GetBooks(ctx context.Context, author string) ([]models.Book, error)
	//GetAuthors returns authors entities by book title from db
	GetAuthors(ctx context.Context, book string) ([]models.Author, error)
}
