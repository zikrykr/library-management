package repository

import (
	"context"

	"github.com/zikrykr/library-management/services/book/config/db"
	"github.com/zikrykr/library-management/services/book/internal/books/model"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type bookStockRepository struct {
	db *db.GormDB
}

func NewBookStockRepository(db *db.GormDB) port.IBookStockRepo {
	return bookStockRepository{db: db}
}

func (r bookStockRepository) CreateBookStock(ctx context.Context, data model.BookStock) error {
	tx, exists := pkg.GetTx(ctx)
	if !exists {
		tx = r.db.DB
	}

	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r bookStockRepository) UpdateBookStock(ctx context.Context, id string, data model.BookStock) error {
	if err := r.db.WithContext(ctx).Model(&model.BookStock{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
