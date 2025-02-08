package service

import (
	"context"
	"errors"

	"github.com/zikrykr/library-management/services/auth/config"
	"github.com/zikrykr/library-management/services/auth/internal/auth/payload"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
	"github.com/zikrykr/library-management/shared/pkg"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repository port.IAuthRepo
}

func NewLoginService(repo port.IAuthRepo) port.ILoginService {
	return LoginService{
		repository: repo,
	}
}

func (s LoginService) Login(ctx context.Context, req payload.LoginReq) (payload.LoginResp, error) {
	var res payload.LoginResp

	userData, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}

	// compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(req.Password)); err != nil {
		return res, errors.New("invalid password")
	}

	// generate JWT
	conf := config.GetConfig()
	jwtClaims := &pkg.JWTClaims{
		UserID: userData.ID,
		Email:  userData.Email,
		Role:   userData.Role,
	}
	jwtConfig := pkg.JWTConfig{
		AppName:   conf.App.Name,
		JWTSecret: conf.App.JWTSecret,
	}
	resToken, err := pkg.GenerateJWT(jwtClaims, jwtConfig)
	if err != nil {
		return res, err
	}

	res.AccessToken = resToken.AccessToken

	return res, nil
}
