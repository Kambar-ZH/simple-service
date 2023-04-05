package managers

import (
	"github.com/Kambar-ZH/simple-service/internal/services/auth_service"
	"github.com/Kambar-ZH/simple-service/internal/services/user_service"
	"sync"
)

var ServiceManager = &Services{}

type Services struct {
	authServiceInit sync.Once
	authService     auth_service.Auth

	userServiceInit sync.Once
	userService     user_service.User
}

func (s *Services) Auth() auth_service.Auth {
	s.authServiceInit.Do(func() {
		s.authService = auth_service.New(
			auth_service.WithUserRepo(repositories.User()),
		)
	})
	return s.authService
}

func (s *Services) User() user_service.User {
	s.userServiceInit.Do(func() {
		s.userService = user_service.New(
			user_service.WithUserRepo(repositories.User()),
		)
	})
	return s.userService
}
