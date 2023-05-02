package user_handler

import (
	"net/http"

	"github.com/Kambar-ZH/simple-service/internal/services/user_service"
	"github.com/Kambar-ZH/simple-service/pkg/tools/auth_tool"
	"github.com/gin-gonic/gin"
)

type User struct {
	userService user_service.User
}

func New(options ...Option) *User {
	ctl := &User{}
	for _, opt := range options {
		opt(ctl)
	}
	return ctl
}

func (ctl *User) Profile(ctx *gin.Context) {
	currentUserID := auth_tool.GetCurrenUserID(ctx)
	ctx.JSON(http.StatusOK, currentUserID)
}
