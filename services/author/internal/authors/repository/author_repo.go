package repository

import (
	"context"

	"github.com/zikrykr/library-management/services/author/config/db"
	"github.com/zikrykr/library-management/services/author/internal/authors/model"
	"github.com/zikrykr/library-management/services/author/internal/authors/payload"
	"github.com/zikrykr/library-management/services/author/internal/authors/port"
	"gorm.io/gorm"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IAuthorRepo {
	return repository{db: db}
}

func (r repository) GetAuthors(ctx context.Context, req payload.GetAuthorsReq) ([]model.Author, int64, error) {
	var (
		res          []model.Author
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

	if err := query.Offset(offset).Limit(req.Limit).Order(req.SortBy).Find(&res).Error; err != nil {
		return res, totalRecords, err
	}

	if err := query.Model(&model.Author{}).Count(&totalRecords).Error; err != nil {
		return res, 0, err
	}

	return res, totalRecords, nil
}

func (r repository) GetAuthorByID(ctx context.Context, id string) (model.Author, error) {
	var res model.Author

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) CreateAuthor(ctx context.Context, data model.Author) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) UpdateAuthor(ctx context.Context, id string, data model.Author) error {
	if err := r.db.WithContext(ctx).Model(&model.Author{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteAuthorByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Author{}).Error; err != nil {
		return err
	}

	return nil
}
