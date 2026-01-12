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
	RegisterFileRouter(r)

	// 静态资源访问: /image/2023-10-27/xxx.jpg
	r.Static("/image", "./tmp/image")
	r.Static("/video", "./tmp/video")
	r.Static("/document", "./tmp/document")
	r.Static("/text", "./tmp/text")
	r.Static("/other", "./tmp/other")

	return r
}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		middleware.Cors(),
	)
}
