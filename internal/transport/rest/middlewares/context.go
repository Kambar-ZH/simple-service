package middlewares

import (
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/pkg/logger"
	"github.com/Kambar-ZH/simple-service/pkg/tools/uuid_tool"
	"github.com/gin-gonic/gin"
)

func SetContextMetadata() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(logger.RequestIDKey, uuid_tool.GetRandomRequestID())
		ctx.Set(logger.AppKey, conf.GlobalConfig.App.Name)
		ctx.Next()
	}
}
