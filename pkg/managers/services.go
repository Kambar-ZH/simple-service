package managers

import (
	"github.com/Kambar-ZH/simple-service/pkg/services/auth_service"
	"sync"
)

var services = &Services{}

type Services struct {
	authServiceInit sync.Once
	authService     auth_service.Auth
}

func (s *Services) Auth() auth_service.Auth {
	s.authServiceInit.Do(func() {
		s.authService = auth_service.New(
			auth_service.WithUserRepo(repositories.User()),
		)
	})
	return s.authService
}
