package auth_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/internal/dtos"
	"github.com/Kambar-ZH/simple-service/internal/models"
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"github.com/Kambar-ZH/simple-service/pkg/tools/auth_tool"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type TestTool struct {
	init     sync.Once
	userRepo user_repo.User
	userSrv  Auth
	rollback func()
}

func (t *TestTool) Init() {
	t.init.Do(func() {
		if err := conf.GlobalConfig.Init(); err != nil {
			panic(err)
		}
		tx := conf.GlobalConfig.GormDB.Begin()

		t.userRepo = user_repo.New(tx)
		t.userSrv = New(
			WithUserRepo(t.userRepo),
			WithLogger(conf.GlobalConfig.Lgr),
		)
		t.rollback = func() { tx.Rollback() }
	})
}

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
	userSrv Auth,
) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		request := dtos.RegisterRequest{
			Email:    a.user.Email,
			Password: auth_tool.Password(a.user.Password),
		}
		_, err := userSrv.Register(ctx, request)
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
	var testTool TestTool
	testTool.Init()
	defer testTool.rollback()

	t.Run("register user", func(t *testing.T) {
		authFlow.register(testTool.userRepo, testTool.userSrv)(t)
	})
}

func (a *AuthFlow) login(
	userSrv Auth,
) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		response, err := userSrv.Login(ctx, dtos.LoginRequest{
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
	var testTool TestTool
	testTool.Init()
	defer testTool.rollback()

	t.Run("login user", func(t *testing.T) {
		authFlow.register(testTool.userRepo, testTool.userSrv)(t)
		authFlow.login(testTool.userSrv)(t)
	})
}

func (a *AuthFlow) refresh(
	userSrv Auth,
) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		response, err := userSrv.Refresh(ctx, dtos.RefreshRequest{
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
	var testTool TestTool
	testTool.Init()
	defer testTool.rollback()

	t.Run("refresh user token", func(t *testing.T) {
		authFlow.register(testTool.userRepo, testTool.userSrv)(t)
		authFlow.login(testTool.userSrv)(t)
		authFlow.refresh(testTool.userSrv)(t)
	})
}
