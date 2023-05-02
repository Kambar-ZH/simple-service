package auth_handler

import (
	"net/http"

	"github.com/Kambar-ZH/simple-service/internal/dtos"
	"github.com/Kambar-ZH/simple-service/internal/services/auth_service"
	"github.com/gin-gonic/gin"
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

// Register
// @Summary      Create user by email and password
// @Description  Create user by email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param 		 payload body dtos.RegisterRequest true "payload"
// @Success      200 {object}  dtos.RegisterResponse
// @Failure      400 {object}  string
// @Failure      500 {object}  string
// @Router       /api/v1/auth/register [post]
func (ctl *Auth) Register(ctx *gin.Context) {
	var payload dtos.RegisterRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
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

// Login
// @Summary      Obtain access and refresh token pair
// @Description  Obtain access and refresh token pair
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param 		 payload body dtos.LoginRequest true "payload"
// @Success      200 {object}  dtos.LoginResponse
// @Failure      400 {object}  string
// @Failure      500 {object}  string
// @Router       /api/v1/auth/login [post]
func (ctl *Auth) Login(ctx *gin.Context) {
	var payload dtos.LoginRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctl.authService.Login(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Refresh
// @Summary      Obtain new token pair by invalidating old refresh token
// @Description  Obtain new token pair by invalidating old refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param 		 payload body dtos.RefreshRequest true "payload"
// @Success      200 {object}  dtos.RefreshResponse
// @Failure      400 {object}  string
// @Failure      500 {object}  string
// @Router       /api/v1/auth/refresh [post]
func (ctl *Auth) Refresh(ctx *gin.Context) {
	var payload dtos.RefreshRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctl.authService.Refresh(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}
