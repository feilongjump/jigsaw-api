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
	r.Static("/image", "./tmp/static/image")
	r.Static("/video", "./tmp/static/video")
	r.Static("/document", "./tmp/static/document")
	r.Static("/text", "./tmp/static/text")
	r.Static("/other", "./tmp/static/other")

	return r
}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		middleware.Cors(),
	)
}
