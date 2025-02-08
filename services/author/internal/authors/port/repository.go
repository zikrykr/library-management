package port

import (
	"context"

	"github.com/zikrykr/library-management/services/author/internal/authors/model"
	"github.com/zikrykr/library-management/services/author/internal/authors/payload"
)

type IAuthorRepo interface {
	CreateAuthor(ctx context.Context, data model.Author) error
	UpdateAuthor(ctx context.Context, id string, data model.Author) error
	GetAuthors(ctx context.Context, req payload.GetAuthorsReq) ([]model.Author, int64, error)
	GetAuthorByID(ctx context.Context, id string) (model.Author, error)
	DeleteAuthorByID(ctx context.Context, id string) error
}
