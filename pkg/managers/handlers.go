package managers

import (
	"github.com/Kambar-ZH/simple-service/pkg/transport/rest/handlers/auth_handler"
	"github.com/Kambar-ZH/simple-service/pkg/transport/rest/handlers/user_handler"
	"sync"
)

var API = &Handlers{}

type Handlers struct {
	authHandlerInit sync.Once
	authHandler     *auth_handler.Auth

	userHandlerInit sync.Once
	userHandler     *user_handler.User
}

func (h *Handlers) Auth() *auth_handler.Auth {
	h.authHandlerInit.Do(func() {
		h.authHandler = auth_handler.New(
			auth_handler.WithAuthService(services.Auth()),
		)
	})
	return h.authHandler
}

func (h *Handlers) User() *user_handler.User {
	h.userHandlerInit.Do(func() {
		h.userHandler = user_handler.New(
			user_handler.WithUserService(services.User()),
		)
	})
	return h.userHandler
}
