package routes

import (
	"github.com/feilongjump/jigsaw-api/app/http/middlewares"
	"github.com/feilongjump/jigsaw-api/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	appConfig := config.GetAppConfig()

	gin.SetMode(appConfig.GinMode)
	r := gin.Default()

	// 注册中间件
	registerMiddlewares(r)

	// 注册路由
	registerRoutes(r)

	return r
}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		middlewares.Cors(),
	)
}
