package user_handler

import (
	"github.com/Kambar-ZH/simple-service/internal/tools/auth_tool"
	"github.com/Kambar-ZH/simple-service/pkg/services/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
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
