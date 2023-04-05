package user_repo

import (
	"context"
	"github.com/Kambar-ZH/simple-service/internal/models"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return &user{db: db}
}

func (u user) GetBy(ctx context.Context, where models.User) (user models.User, err error) {

	err = u.db.
		Table(models.User{}.TableName()).
		Where(where).
		First(&user).
		Error

	if err != nil {
		return
	}

	return
}

func (u user) Save(ctx context.Context, model models.User) (result models.User, err error) {

	err = u.db.Table(models.User{}.TableName()).
		Save(&model).
		Error

	if err != nil {
		return
	}

	result = model

	return
}
