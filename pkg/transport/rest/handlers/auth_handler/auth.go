package auth_handler

import (
	"github.com/Kambar-ZH/simple-service/pkg/dtos"
	"github.com/Kambar-ZH/simple-service/pkg/services/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	authService auth_service.Auth
}

func New(options ...Option) *Auth {
	ctl := &Auth{}
	for _, opt := range options {
		opt(ctl)
	}
	return ctl
}

func (ctl *Auth) Register(ctx *gin.Context) {
	var (
		payload dtos.RegisterRequest
	)
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := ctl.authService.Register(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (ctl *Auth) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func (ctl *Auth) Refresh(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func (ctl *Auth) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "ok"})
}
