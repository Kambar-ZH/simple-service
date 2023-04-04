package user_service

import (
	"context"
	"github.com/Kambar-ZH/simple-service/pkg/models"
)

type User interface {
	GetBy(ctx context.Context, where models.User) (result models.User, err error)
}
