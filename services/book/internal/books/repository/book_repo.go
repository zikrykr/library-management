package repository

import (
	"context"

	"github.com/zikrykr/library-management/services/book/config/db"
	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"gorm.io/gorm"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IBookRepo {
	return repository{db: db}
}

func (r repository) GetBooks(ctx context.Context, req payload.GetBooksReq) ([]model.Book, int64, error) {
	var (
		res          []model.Book
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

	if req.Title != "" {
		// use ts_query for full text search and add ILIKE for case insensitive search and incomplete words
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("to_tsvector('english', title) @@ plainto_tsquery(?) OR title ILIKE ?", req.Title, "%"+req.Title+"%")
		})
	}

	if req.AuthorID != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("author_id = ?", req.AuthorID)
		})
	}

	if req.CategoryID != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("category_id = ?", req.CategoryID)
		})
	}

	if req.SortBy == "" {
		req.SortBy = "created_at DESC"
	}

	query := r.db.WithContext(ctx).Scopes(fScopes...)

	if err := query.Offset(offset).Limit(req.Limit).Order(req.SortBy).Find(&res).Error; err != nil {
		return res, totalRecords, err
	}

	if err := query.Model(&model.Book{}).Count(&totalRecords).Error; err != nil {
		return res, 0, err
	}

	return res, totalRecords, nil
}

func (r repository) GetBookByID(ctx context.Context, id string) (model.Book, error) {
	var res model.Book

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) CreateBook(ctx context.Context, data model.Book) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) UpdateBook(ctx context.Context, id string, data model.Book) error {
	if err := r.db.WithContext(ctx).Model(&model.Book{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteBookByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Book{}).Error; err != nil {
		return err
	}

	return nil
}
