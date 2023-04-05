package auth_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/internal/dtos"
	"github.com/Kambar-ZH/simple-service/internal/models"
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"github.com/Kambar-ZH/simple-service/pkg/tools/auth_tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

type AuthFlow struct {
	user   models.User
	tokens dtos.TokenPair
}

var authFlow = AuthFlow{
	user: models.User{
		Email:    "test",
		Password: "test",
	},
}

func (a *AuthFlow) register(
	userRepo user_repo.User,
	srv Auth,
) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		request := dtos.RegisterRequest{
			Email:    a.user.Email,
			Password: auth_tool.Password(a.user.Password),
		}
		_, err := srv.Register(ctx, request)
		if err != nil {
			t.Error(err)
		}
		user, err := userRepo.GetBy(ctx, models.User{Email: request.Email})
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, request.Email, user.Email, "emails should be equal")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestAuth_Register(t *testing.T) {
	if err := conf.GlobalConfig.Init(); err != nil {
		panic(err)
	}
	tx := conf.GlobalConfig.GormDB.Begin()
	defer tx.Rollback()

	userRepo := user_repo.New(tx)
	srv := New(WithUserRepo(userRepo))

	t.Run("register user", func(t *testing.T) {
		authFlow.register(userRepo, srv)(t)
	})
}

func (a *AuthFlow) login(
	srv Auth,
) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		response, err := srv.Login(ctx, dtos.LoginRequest{
			Email:    a.user.Email,
			Password: auth_tool.Password(a.user.Password),
		})
		if err != nil {
			t.Error(err)
		}
		a.tokens = response.TokenPair
		assert.NotEmpty(t, response.Access, "access token should not be empty")
		assert.NotEmpty(t, response.Refresh, "refresh token should not be empty")
	}
}

func TestAuth_Login(t *testing.T) {
	if err := conf.GlobalConfig.Init(); err != nil {
		panic(err)
	}
	tx := conf.GlobalConfig.GormDB.Begin()
	defer tx.Rollback()

	userRepo := user_repo.New(tx)
	srv := New(WithUserRepo(userRepo))

	t.Run("login user", func(t *testing.T) {
		authFlow.register(userRepo, srv)(t)
		authFlow.login(srv)(t)
	})
}

func (a *AuthFlow) refresh(
	srv Auth,
) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		response, err := srv.Refresh(ctx, dtos.RefreshRequest{
			Refresh: a.tokens.Refresh,
		})
		if err != nil {
			t.Error(err)
		}
		assert.NotEmpty(t, response.Access, "access token should not be empty")
		assert.NotEmpty(t, response.Refresh, "refresh token should not be empty")
	}
}

func TestAuth_Refresh(t *testing.T) {
	if err := conf.GlobalConfig.Init(); err != nil {
		panic(err)
	}
	tx := conf.GlobalConfig.GormDB.Begin()
	defer tx.Rollback()

	userRepo := user_repo.New(tx)
	srv := New(WithUserRepo(userRepo))

	t.Run("refresh user token", func(t *testing.T) {
		authFlow.register(userRepo, srv)(t)
		authFlow.login(srv)(t)
		authFlow.refresh(srv)(t)
	})
}
