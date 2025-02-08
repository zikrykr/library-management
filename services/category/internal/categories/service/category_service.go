package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/zikrykr/library-management/services/category/internal/categories/model"
	"github.com/zikrykr/library-management/services/category/internal/categories/payload"
	"github.com/zikrykr/library-management/services/category/internal/categories/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type CategoryService struct {
	repository port.ICategoryRepo
}

func NewCategoryService(repo port.ICategoryRepo) port.ICategoryService {
	return CategoryService{
		repository: repo,
	}
}

func (s CategoryService) CreateCategory(ctx context.Context, req payload.CreateCategoryReq) error {
	id := uuid.New()

	data := model.Category{
		ID:          id.String(),
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repository.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}

func (s CategoryService) UpdateCategory(ctx context.Context, id string, req payload.UpdateCategoryReq) error {
	data := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repository.UpdateCategory(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (s CategoryService) GetCategories(ctx context.Context, req payload.GetCategoriesReq) ([]model.Category, *pkg.Pagination, error) {
	req.Limit = pkg.ValidateLimit(req.Limit)
	req.Page = pkg.ValidatePage(req.Page)

	authors, totalRecords, err := s.repository.GetCategories(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	pagination := pkg.CalculatePagination(totalRecords, int64(req.Page), int64(req.Limit), req.SortBy)

	return authors, pagination, nil
}

func (s CategoryService) GetCategoryByID(ctx context.Context, id string) (model.Category, error) {
	author, err := s.repository.GetCategoryByID(ctx, id)
	if err != nil {
		return model.Category{}, err
	}

	return author, nil
}

func (s CategoryService) DeleteCategoryByID(ctx context.Context, id string) error {
	if err := s.repository.DeleteCategoryByID(ctx, id); err != nil {
		return err
	}

	return nil
}
