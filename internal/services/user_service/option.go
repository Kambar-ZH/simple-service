package user_service

import "github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"

type Option func(srv *user)

func WithUserRepo(userRepo user_repo.User) Option {
	return func(srv *user) {
		srv.userRepo = userRepo
	}
}
