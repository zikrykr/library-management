package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"github.com/zikrykr/library-management/shared/pkg"
	"gorm.io/gorm"
)

type BookService struct {
	DB                  *gorm.DB
	bookRepository      port.IBookRepo
	bookStockRepository port.IBookStockRepo
}

func NewBookService(db *gorm.DB, bookRepo port.IBookRepo, bookStockRepo port.IBookStockRepo) port.IBookService {
	return BookService{
		DB:                  db,
		bookRepository:      bookRepo,
		bookStockRepository: bookStockRepo,
	}
}

func (s BookService) CreateBook(ctx context.Context, req payload.CreateBookReq) error {
	bookID := uuid.New().String()

	data := model.Book{
		ID:            bookID,
		Title:         req.Title,
		Description:   req.Description,
		ISBN:          req.ISBN,
		AuthorID:      req.AuthorID,
		CategoryID:    req.CategoryID,
		PublishedYear: req.PublishedYear,
	}

	stockData := model.BookStock{
		ID:             uuid.New().String(),
		BookID:         bookID,
		TotalStock:     req.TotalStock,
		AvailableStock: req.AvailableStock,
	}

	// use transaction to create book and book stock
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		ctx := pkg.SetTx(ctx, tx)

		if err := s.bookRepository.CreateBook(ctx, data); err != nil {
			return err
		}

		if err := s.bookStockRepository.CreateBookStock(ctx, stockData); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil

}

func (s BookService) UpdateBook(ctx context.Context, id string, req payload.UpdateBookReq) error {
	data := model.Book{
		Title:         req.Title,
		Description:   req.Description,
		ISBN:          req.ISBN,
		AuthorID:      req.AuthorID,
		CategoryID:    req.CategoryID,
		PublishedYear: req.PublishedYear,
	}

	if err := s.bookRepository.UpdateBook(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (s BookService) GetBooks(ctx context.Context, req payload.GetBooksReq) ([]model.Book, *pkg.Pagination, error) {
	req.Limit = pkg.ValidateLimit(req.Limit)
	req.Page = pkg.ValidatePage(req.Page)

	books, totalRecords, err := s.bookRepository.GetBooks(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	pagination := pkg.CalculatePagination(totalRecords, int64(req.Page), int64(req.Limit), req.SortBy)

	return books, pagination, nil
}

func (s BookService) GetBookByID(ctx context.Context, id string) (model.Book, error) {
	book, err := s.bookRepository.GetBookByID(ctx, id)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (s BookService) DeleteBookByID(ctx context.Context, id string) error {
	if err := s.bookRepository.DeleteBookByID(ctx, id); err != nil {
		return err
	}

	return nil
}
