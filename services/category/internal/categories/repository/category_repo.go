package repository

import (
	"context"

	"github.com/zikrykr/library-management/services/category/config/db"
	"github.com/zikrykr/library-management/services/category/internal/categories/model"
	"github.com/zikrykr/library-management/services/category/internal/categories/payload"
	"github.com/zikrykr/library-management/services/category/internal/categories/port"
	"gorm.io/gorm"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.ICategoryRepo {
	return repository{db: db}
}

func (r repository) GetCategories(ctx context.Context, req payload.GetCategoriesReq) ([]model.Category, int64, error) {
	var (
		res          []model.Category
		fScopes      []func(db *gorm.DB) *gorm.DB
		totalRecords int64
	)

	offset := (req.Page - 1) * req.Limit

	if len(req.IDIN) > 0 {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id NOT IN (?)", req.IDIN)
		})
	}

	if req.ID != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id != ?", req.ID)
		})
	}

	if req.Name != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name != ?", req.Name)
		})
	}

	if req.SortBy == "" {
		req.SortBy = "created_at DESC"
	}

	query := r.db.WithContext(ctx).Scopes(fScopes...)

	if err := query.Model(&model.Category{}).Count(&totalRecords).Error; err != nil {
		return res, 0, err
	}

	if err := query.Offset(offset).Limit(req.Limit).Order(req.SortBy).Find(&res).Error; err != nil {
		return res, totalRecords, err
	}

	return res, totalRecords, nil
}

func (r repository) GetCategoryByID(ctx context.Context, id string) (model.Category, error) {
	var res model.Category

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) CreateCategory(ctx context.Context, data model.Category) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) UpdateCategory(ctx context.Context, id string, data model.Category) error {
	if err := r.db.WithContext(ctx).Model(&model.Category{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteCategoryByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return err
	}

	return nil
}
