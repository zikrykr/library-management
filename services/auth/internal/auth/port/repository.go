package port

import (
	"context"

	"github.com/zikrykr/library-management/services/auth/internal/auth/model"
)

type IAuthRepo interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, data model.User) error
}
