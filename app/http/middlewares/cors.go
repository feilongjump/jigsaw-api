package middlewares

import (
	"github.com/feilongjump/jigsaw-api/app/http/responses"
	"github.com/gin-gonic/gin"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		if method != "" {
			// 可将 * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		}

		if method == "OPTIONS" {
			responses.SuccessNoContent(ctx)
			return
		}

		ctx.Next()
	}
}
