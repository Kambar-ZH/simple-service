package rest

import (
	"github.com/Kambar-ZH/simple-service/pkg/managers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()

	api := router.Group("/api")

	v1 := api.Group("/v1")

	auth := v1.Group("/auth_service")
	{
		auth.GET("/login", managers.API.Auth().Login)
		auth.GET("/register", managers.API.Auth().Register)
		auth.GET("/refresh", managers.API.Auth().Refresh)
		auth.GET("/logout", managers.API.Auth().Logout)
	}

	router.Run("127.0.0.1:8000")
}
