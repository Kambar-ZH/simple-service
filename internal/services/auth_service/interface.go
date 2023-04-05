package auth_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/internal/dtos"
)

type Auth interface {
	Register(ctx context.Context, request dtos.RegisterRequest) (result dtos.RegisterResponse, err error)
	Login(ctx context.Context, request dtos.LoginRequest) (result dtos.LoginResponse, err error)
	Refresh(ctx context.Context, request dtos.RefreshRequest) (result dtos.RefreshResponse, err error)
}
