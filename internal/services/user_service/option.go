package user_service

import (
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"github.com/Kambar-ZH/simple-service/pkg/logger"
)

type Option func(srv *user)

func WithUserRepo(userRepo user_repo.User) Option {
	return func(srv *user) {
		srv.userRepo = userRepo
	}
}

func WithLogger(lgr logger.Logger) Option {
	return func(srv *user) {
		srv.lgr = lgr
	}
}
