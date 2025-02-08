package port

import (
	"context"

	"github.com/zikrykr/library-management/services/category/internal/categories/model"
	"github.com/zikrykr/library-management/services/category/internal/categories/payload"
	"github.com/zikrykr/library-management/shared/pkg"
)

type ICategoryService interface {
	CreateCategory(ctx context.Context, req payload.CreateCategoryReq) error
	UpdateCategory(ctx context.Context, id string, req payload.UpdateCategoryReq) error
	GetCategories(ctx context.Context, req payload.GetCategoriesReq) ([]model.Category, *pkg.Pagination, error)
	GetCategoryByID(ctx context.Context, id string) (model.Category, error)
	DeleteCategoryByID(ctx context.Context, id string) error
}
