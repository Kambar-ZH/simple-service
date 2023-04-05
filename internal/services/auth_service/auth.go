package auth_service

import (
	"context"
	"errors"
	"github.com/Kambar-ZH/simple-service/internal/dtos"
	"github.com/Kambar-ZH/simple-service/internal/models"
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"github.com/Kambar-ZH/simple-service/pkg/tools/auth_tool"
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

var ErrInvalidPasswordOrEmail = errors.New("invalid password or email")
var ErrTokenIsInvalid = errors.New("token is invalid")

func (srv *auth) Register(ctx context.Context, request dtos.RegisterRequest) (result dtos.RegisterResponse, err error) {
	hashedPassword, err := request.Password.Hash()
	if err != nil {
		return
	}

	model, err := srv.userRepo.Save(ctx, models.User{
		Email:    request.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return
	}

	result.ID = model.ID

	return
}

func (srv *auth) Login(ctx context.Context, request dtos.LoginRequest) (result dtos.LoginResponse, err error) {
	user, err := srv.userRepo.GetBy(ctx, models.User{Email: request.Email})
	if err != nil {
		err = ErrInvalidPasswordOrEmail
		return
	}

	if ok := auth_tool.CheckPasswordHash(string(request.Password), user.Password); !ok {
		err = ErrInvalidPasswordOrEmail
		return
	}

	access, refresh, err := auth_tool.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		return
	}

	return dtos.LoginResponse{
		TokenPair: dtos.TokenPair{
			Access:  access,
			Refresh: refresh,
		},
	}, nil
}

func (srv *auth) Refresh(ctx context.Context, request dtos.RefreshRequest) (result dtos.RefreshResponse, err error) {
	token, err := auth_tool.ParseToken(request.Refresh)
	if err != nil {
		return
	}
	if !token.Valid {
		err = ErrTokenIsInvalid
		return
	}

	user, err := srv.userRepo.GetBy(ctx, models.User{
		ID: auth_tool.GetCurrenUserIDByToken(token),
	})
	if err != nil {
		return
	}

	access, refresh, err := auth_tool.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		return
	}

	return dtos.RefreshResponse{
		TokenPair: dtos.TokenPair{
			Access:  access,
			Refresh: refresh,
		},
	}, nil
}
