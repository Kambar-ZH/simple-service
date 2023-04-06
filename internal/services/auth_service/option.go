package auth_service

import (
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"github.com/Kambar-ZH/simple-service/pkg/logger"
)

type Option func(srv *auth)

func WithUserRepo(userRepo user_repo.User) Option {
	return func(srv *auth) {
		srv.userRepo = userRepo
	}
}

func WithLogger(lgr logger.Logger) Option {
	return func(srv *auth) {
		srv.lgr = lgr
	}
}
