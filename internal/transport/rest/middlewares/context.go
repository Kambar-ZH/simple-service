package middlewares

import (
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/pkg/logger"
	"github.com/Kambar-ZH/simple-service/pkg/tools/tracing_tools"
	"github.com/Kambar-ZH/simple-service/pkg/tools/uuid_tool"
	"github.com/gin-gonic/gin"
)

func SetContextMetadata() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tm := tracing_tools.TracingMetadata{
			RequestID: uuid_tool.GetRandomRequestID(),
			AppName:   conf.GlobalConfig.App.Name,
		}
		ctx.Set(logger.TracingMetadataKey, tm)
	}
}
