package handlers

import (
	"sync"

	"github.com/Kambar-ZH/simple-service/internal/managers"
	"github.com/Kambar-ZH/simple-service/internal/transport/rest/handlers/auth_handler"
	"github.com/Kambar-ZH/simple-service/internal/transport/rest/handlers/user_handler"
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
			auth_handler.WithAuthService(managers.ServiceManager.Auth()),
		)
	})
	return h.authHandler
}

func (h *Handlers) User() *user_handler.User {
	h.userHandlerInit.Do(func() {
		h.userHandler = user_handler.New(
			user_handler.WithUserService(managers.ServiceManager.User()),
		)
	})
	return h.userHandler
}
