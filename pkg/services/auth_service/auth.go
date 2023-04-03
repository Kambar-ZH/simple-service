package auth_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/pkg/dtos"
	"github.com/Kambar-ZH/simple-service/pkg/models"
	"github.com/Kambar-ZH/simple-service/pkg/repositories/common/user_repo"
)

type auth struct {
	userRepo user_repo.User
}

func New(options ...Option) Auth {
	srv := &auth{}
	for _, opt := range options {
		opt(srv)
	}
	return srv
}

func (srv auth) Register(ctx context.Context, request dtos.RegisterRequest) (result dtos.RegisterResponse, err error) {
	model, err := srv.userRepo.Save(ctx, models.User{Email: request.Email, Password: request.Password})
	if err != nil {
		return
	}

	result.ID = model.ID

	return
}

func (srv auth) Login(ctx context.Context, request dtos.LoginRequest) (err error) {

	return
}

func (srv auth) Refresh(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

func (srv auth) Logout(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}
