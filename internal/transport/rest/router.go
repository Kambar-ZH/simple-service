package rest

import (
	_ "github.com/Kambar-ZH/simple-service/docs"
	"github.com/Kambar-ZH/simple-service/internal/transport/rest/handlers"
	"github.com/Kambar-ZH/simple-service/internal/transport/rest/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func InitRouter() (err error) {
	router := gin.New()

	api := router.Group("/api")

	v1 := api.Group("/v1", middlewares.SetContextMetadata())

	auth := v1.Group("/auth")
	{
		auth.POST("/login", handlers.API.Auth().Login)
		auth.POST("/register", handlers.API.Auth().Register)
		auth.POST("/refresh", handlers.API.Auth().Refresh)
	}

	user := v1.Group("/user")
	{
		user.GET("/profile", middlewares.Authenticated(), handlers.API.User().Profile)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router.Run("0.0.0.0:8000")
}
