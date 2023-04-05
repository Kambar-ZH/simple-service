package auth_handler

import "github.com/Kambar-ZH/simple-service/internal/services/auth_service"

type Option func(ctl *Auth)

func WithAuthService(authService auth_service.Auth) Option {
	return func(ctl *Auth) {
		ctl.authService = authService
	}
}
