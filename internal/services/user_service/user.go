package user_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/internal/models"
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"github.com/Kambar-ZH/simple-service/pkg/logger"
)

type user struct {
	userRepo user_repo.User
	lgr      logger.Logger
}

func New(options ...Option) User {
	srv := &user{}
	for _, opt := range options {
		opt(srv)
	}
	return srv
}

func (u user) GetBy(ctx context.Context, where models.User) (result models.User, err error) {
	return u.userRepo.GetBy(ctx, where)
}
