package repository

import (
	"context"

	"github.com/zikrykr/library-management/services/auth/config/db"
	"github.com/zikrykr/library-management/services/auth/internal/auth/model"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IAuthRepo {
	return repository{db: db}
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var res model.User

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) CreateUser(ctx context.Context, data model.User) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
