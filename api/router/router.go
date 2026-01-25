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
	RegisterLedgerCategoryRouter(r)
	RegisterUserWalletRouter(r)
	RegisterLedgerRecordRouter(r)

	// 静态资源访问: /image/2023-10-27/xxx.jpg
	r.Static("/image", "./static/image")
	r.Static("/video", "./static/video")
	r.Static("/document", "./static/document")
	r.Static("/text", "./static/text")
	r.Static("/other", "./static/other")

	return r
}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		middleware.Cors(),
	)
}
