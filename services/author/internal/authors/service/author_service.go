package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/zikrykr/library-management/services/author/internal/authors/model"
	"github.com/zikrykr/library-management/services/author/internal/authors/payload"
	"github.com/zikrykr/library-management/services/author/internal/authors/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type AuthorService struct {
	repository port.IAuthorRepo
}

func NewAuthorService(repo port.IAuthorRepo) port.IAuthorService {
	return AuthorService{
		repository: repo,
	}
}

func (s AuthorService) CreateAuthor(ctx context.Context, req payload.CreateAuthorReq) error {
	id := uuid.New()

	data := model.Author{
		ID:   id.String(),
		Name: req.Name,
		Bio:  req.Bio,
	}

	if err := s.repository.CreateAuthor(ctx, data); err != nil {
		return err
	}

	return nil
}

func (s AuthorService) UpdateAuthor(ctx context.Context, id string, req payload.UpdateAuthorReq) error {
	data := model.Author{
		Name: req.Name,
		Bio:  req.Bio,
	}

	if err := s.repository.UpdateAuthor(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (s AuthorService) GetAuthors(ctx context.Context, req payload.GetAuthorsReq) ([]model.Author, *pkg.Pagination, error) {
	req.Limit = pkg.ValidateLimit(req.Limit)
	req.Page = pkg.ValidatePage(req.Page)

	authors, totalRecords, err := s.repository.GetAuthors(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	pagination := pkg.CalculatePagination(totalRecords, int64(req.Page), int64(req.Limit), req.SortBy)

	return authors, pagination, nil
}

func (s AuthorService) GetAuthorByID(ctx context.Context, id string) (model.Author, error) {
	author, err := s.repository.GetAuthorByID(ctx, id)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (s AuthorService) DeleteAuthorByID(ctx context.Context, id string) error {
	if err := s.repository.DeleteAuthorByID(ctx, id); err != nil {
		return err
	}

	return nil
}
