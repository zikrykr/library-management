package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type BookService struct {
	repository port.IBookRepo
}

func NewBookService(repo port.IBookRepo) port.IBookService {
	return BookService{
		repository: repo,
	}
}

func (s BookService) CreateBook(ctx context.Context, req payload.CreateBookReq) error {
	id := uuid.New()

	data := model.Book{
		ID:            id.String(),
		Title:         req.Title,
		Description:   req.Description,
		ISBN:          req.ISBN,
		AuthorID:      req.AuthorID,
		CategoryID:    req.CategoryID,
		PublishedYear: req.PublishedYear,
	}

	if err := s.repository.CreateBook(ctx, data); err != nil {
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

	if err := s.repository.UpdateBook(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (s BookService) GetBooks(ctx context.Context, req payload.GetBooksReq) ([]model.Book, *pkg.Pagination, error) {
	req.Limit = pkg.ValidateLimit(req.Limit)
	req.Page = pkg.ValidatePage(req.Page)

	books, totalRecords, err := s.repository.GetBooks(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	pagination := pkg.CalculatePagination(totalRecords, int64(req.Page), int64(req.Limit), req.SortBy)

	return books, pagination, nil
}

func (s BookService) GetBookByID(ctx context.Context, id string) (model.Book, error) {
	book, err := s.repository.GetBookByID(ctx, id)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (s BookService) DeleteBookByID(ctx context.Context, id string) error {
	if err := s.repository.DeleteBookByID(ctx, id); err != nil {
		return err
	}

	return nil
}
