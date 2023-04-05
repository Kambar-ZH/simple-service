package user_repo

import (
	"context"
	"github.com/Kambar-ZH/simple-service/internal/models"
)

type User interface {
	GetBy(ctx context.Context, where models.User) (user models.User, err error)
	Save(ctx context.Context, model models.User) (result models.User, err error)
}
