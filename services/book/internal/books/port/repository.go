package port

import (
	"context"

	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
)

type IBookRepo interface {
	CreateBook(ctx context.Context, data model.Book) error
	UpdateBook(ctx context.Context, id string, data model.Book) error
	GetBooks(ctx context.Context, req payload.GetBooksReq) ([]model.Book, int64, error)
	GetBookByID(ctx context.Context, id string) (model.Book, error)
	DeleteBookByID(ctx context.Context, id string) error
}

type IBookStockRepo interface {
	CreateBookStock(ctx context.Context, data model.BookStock) error
	UpdateBookStockByBookID(ctx context.Context, bookID string, data model.BookStock) error
}
