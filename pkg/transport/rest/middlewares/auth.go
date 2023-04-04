package middlewares

import (
	"errors"
	"github.com/Kambar-ZH/simple-service/internal/tools/auth_tool"
	"github.com/gin-gonic/gin"
	"net/http"
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
		auth_tool.SetToken(ctx, token)
		ctx.Next()
	}
}
