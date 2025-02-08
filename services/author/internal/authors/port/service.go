package port

import (
	"context"

	"github.com/zikrykr/library-management/services/author/internal/authors/model"
	"github.com/zikrykr/library-management/services/author/internal/authors/payload"
	"github.com/zikrykr/library-management/shared/pkg"
)

type IAuthorService interface {
	CreateAuthor(ctx context.Context, req payload.CreateAuthorReq) error
	UpdateAuthor(ctx context.Context, id string, req payload.UpdateAuthorReq) error
	GetAuthors(ctx context.Context, req payload.GetAuthorsReq) ([]model.Author, *pkg.Pagination, error)
	GetAuthorByID(ctx context.Context, id string) (model.Author, error)
	DeleteAuthorByID(ctx context.Context, id string) error
}
