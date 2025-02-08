package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zikrykr/library-management/services/auth/internal/auth/model"
	"github.com/zikrykr/library-management/services/auth/internal/auth/payload"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpService struct {
	repository port.IAuthRepo
}

func NewSignUpService(repo port.IAuthRepo) port.ISignUpService {
	return SignUpService{
		repository: repo,
	}
}

func (s SignUpService) SignUp(ctx context.Context, req payload.SignUpReq) error {
	// check if user has been registered
	if err := s.checkUserRegistered(ctx, req.Email); err != nil {
		return err
	}

	id := uuid.New()

	data := model.User{
		ID:       id.String(),
		FullName: req.FullName,
		Email:    req.Email,
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	data.PasswordHash = string(hashedPw)

	if err := s.repository.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}

func (s SignUpService) checkUserRegistered(ctx context.Context, email string) error {
	// check if email has been registered
	resEmail, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if resEmail.ID != "" {
		return errors.New("email has been registered")
	}

	return nil
}
