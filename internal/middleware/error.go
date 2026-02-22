package middleware

import (
	"net/http"

	"jigsaw-api/pkg/logger"
	"jigsaw-api/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.Log.Error("请求错误", zap.String("error", e.Error()))
			}

			if !c.Writer.Written() {
				response.ErrorWithStatus(c, http.StatusInternalServerError, 500, "服务器内部错误")
			}
		}
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("Panic 已恢复", zap.Any("error", err))
				response.ErrorWithStatus(c, http.StatusInternalServerError, 500, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
