package service

import (
	"context"

	"github.com/zikrykr/library-management/services/auth/internal/auth/payload"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
)

type ProfileService struct {
	repository port.IAuthRepo
}

func NewProfileService(repo port.IAuthRepo) port.IProfileService {
	return ProfileService{
		repository: repo,
	}
}

func (s ProfileService) GetProfile(ctx context.Context, userEmail string) (payload.GetProfileResp, error) {
	userData, err := s.repository.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return payload.GetProfileResp{}, err
	}

	res := payload.GetProfileResp{
		ID:       userData.ID,
		FullName: userData.FullName,
		Email:    userData.Email,
		Role:     userData.Role,
	}

	return res, nil
}
