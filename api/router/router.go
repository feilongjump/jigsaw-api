package router

import (
	"github.com/feilongjump/jigsaw-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	r := gin.Default()

	// 注册中间件
	registerMiddlewares(r)

	RegisterNoteRouter(r)
	RegisterUserRouter(r)

	return r
}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		middleware.Cors(),
	)
}
