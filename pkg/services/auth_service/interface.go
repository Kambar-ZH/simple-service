package auth_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/pkg/dtos"
)

type Auth interface {
	Register(ctx context.Context, request dtos.RegisterRequest) (result dtos.RegisterResponse, err error)
	Login(ctx context.Context, request dtos.LoginRequest) (err error)
	Refresh(ctx context.Context) (err error)
	Logout(ctx context.Context) (err error)
}
