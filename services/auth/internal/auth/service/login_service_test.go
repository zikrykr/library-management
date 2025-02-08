package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zikrykr/library-management/services/auth/internal/auth/model"
	"github.com/zikrykr/library-management/services/auth/internal/auth/payload"
	"github.com/zikrykr/library-management/services/auth/mock"
)

func TestLoginService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		mockAuthRepo = mock.NewMockIAuthRepo(ctrl)
	)

	type args struct {
		ctx context.Context
		req payload.LoginReq
	}
	tests := []struct {
		name        string
		args        args
		mockCallsFn []*gomock.Call
		wantRes     payload.LoginResp
		wantErr     bool
	}{
		{
			name: "Successfully Login",
			args: args{
				ctx: context.Background(),
				req: payload.LoginReq{
					Email:    "email@email.com",
					Password: "jikur123",
				},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{
					PasswordHash: "password_hashed",
				}, nil),
			},
			wantRes: payload.LoginResp{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginService := &LoginService{
				repository: mockAuthRepo,
			}

			gotRes, err := loginService.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginService.Login() error=%v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("loginService.Login gotRes=%v want %v", gotRes, tt.wantRes)
			}
		})
	}
}
