package middlewares

import (
	"errors"
	"net/http"

	"github.com/Kambar-ZH/simple-service/pkg/tools/auth_tool"
	"github.com/gin-gonic/gin"
)

var ErrUnauthorized = errors.New("unauthorized request")

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := auth_tool.ParseToken(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		}
		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}
		auth_tool.SetAccessToken(ctx, token)
	}
}
