package port

import (
	"context"

	"github.com/zikrykr/library-management/services/category/internal/categories/model"
	"github.com/zikrykr/library-management/services/category/internal/categories/payload"
)

type ICategoryRepo interface {
	CreateCategory(ctx context.Context, data model.Category) error
	UpdateCategory(ctx context.Context, id string, data model.Category) error
	GetCategories(ctx context.Context, req payload.GetCategoriesReq) ([]model.Category, int64, error)
	GetCategoryByID(ctx context.Context, id string) (model.Category, error)
	DeleteCategoryByID(ctx context.Context, id string) error
}
