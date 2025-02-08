package port

import (
	"context"

	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
	"github.com/zikrykr/library-management/shared/pkg"
)

type IBookService interface {
	CreateBook(ctx context.Context, req payload.CreateBookReq) error
	UpdateBook(ctx context.Context, id string, req payload.UpdateBookReq) error
	GetBooks(ctx context.Context, req payload.GetBooksReq) ([]model.Book, *pkg.Pagination, error)
	GetBookByID(ctx context.Context, id string) (model.Book, error)
	DeleteBookByID(ctx context.Context, id string) error
}
