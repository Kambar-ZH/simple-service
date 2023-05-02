package user_repo

import (
	"context"
	"errors"
	"testing"

	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser_Save(t *testing.T) {
	if err := conf.GlobalConfig.Init(); err != nil {
		panic(err)
	}
	tx := conf.GlobalConfig.GormDB.Begin()
	defer tx.Rollback()

	t.Run("insert user", func(t *testing.T) {
		repo := New(tx)
		ctx := context.Background()
		model := models.User{
			Name:     "petr",
			Surname:  "bob",
			Email:    "petr@mail.ru",
			Password: "pass",
		}
		insertResult, err := repo.Save(ctx, model)
		assert.Equal(t, nil, err, "error should be nil")
		selectResult, err := repo.GetBy(ctx, models.User{ID: insertResult.ID})
		assert.Equal(t, nil, err, "error should be nil")
		assert.Equal(t, insertResult, selectResult, "insert and select should be equal")
	})
}

func TestUser_GetBy(t *testing.T) {
	if err := conf.GlobalConfig.Init(); err != nil {
		panic(err)
	}
	tx := conf.GlobalConfig.GormDB.Begin()
	defer tx.Rollback()

	t.Run("get user by name, not found error", func(t *testing.T) {
		repo := New(tx)
		ctx := context.Background()
		where := models.User{
			Name: "petr",
		}
		_, err := repo.GetBy(ctx, where)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Error("error should be: record not found")
		}
	})

	t.Run("get user by name", func(t *testing.T) {
		repo := New(tx)
		ctx := context.Background()
		where := models.User{
			Email: "kambar@mail.ru",
		}
		result, err := repo.GetBy(ctx, where)
		if err != nil {
			t.Errorf("error should be nil")
		}
		assert.Equal(t, where.Name, result.Name, "emails should be equal")
	})
}
