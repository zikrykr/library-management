package repository

import (
	"context"

	"github.com/zikrykr/library-management/services/book/config/db"
	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"github.com/zikrykr/library-management/shared/pkg"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *db.GormDB
}

func NewBookRepository(db *db.GormDB) port.IBookRepo {
	return bookRepository{db: db}
}

func (r bookRepository) GetBooks(ctx context.Context, req payload.GetBooksReq) ([]model.Book, int64, error) {
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

	if err := query.Model(&model.Book{}).Count(&totalRecords).Error; err != nil {
		return res, 0, err
	}

	if err := query.Offset(offset).Limit(req.Limit).Order(req.SortBy).Find(&res).Error; err != nil {
		return res, totalRecords, err
	}

	return res, totalRecords, nil
}

func (r bookRepository) GetBookByID(ctx context.Context, id string) (model.Book, error) {
	var res model.Book

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r bookRepository) CreateBook(ctx context.Context, data model.Book) error {
	tx, exists := pkg.GetTx(ctx)
	if !exists {
		tx = r.db.DB
	}

	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r bookRepository) UpdateBook(ctx context.Context, id string, data model.Book) error {
	tx, exists := pkg.GetTx(ctx)
	if !exists {
		tx = r.db.DB
	}

	if err := tx.WithContext(ctx).Model(&model.Book{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (r bookRepository) DeleteBookByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Book{}).Error; err != nil {
		return err
	}

	return nil
}
