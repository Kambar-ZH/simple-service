package managers

import (
	"github.com/Kambar-ZH/simple-service/pkg/transport/rest/handlers/auth_handler"
	"sync"
)

var API = &Handlers{}

type Handlers struct {
	authHandlerInit sync.Once
	authHandler     *auth_handler.Auth
}

func (h *Handlers) Auth() *auth_handler.Auth {
	h.authHandlerInit.Do(func() {
		h.authHandler = auth_handler.New(
			auth_handler.WithAuthService(services.Auth()),
		)
	})
	return h.authHandler
}
