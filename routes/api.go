package routes

import (
	"github.com/feilongjump/jigsaw-api/app/http/controllers"
	"github.com/feilongjump/jigsaw-api/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {

	baseRoutes(r)

	authRoutes(r)

	userRoutes(r)
}

func baseRoutes(r *gin.Engine) {
	homeC := new(controllers.HomeController)
	r.GET("/", homeC.Index)
}

func authRoutes(r *gin.Engine) {
	authC := new(controllers.AuthController)

	authG := r.Group("/auth")
	authG.POST("login", authC.Login)
	authG.POST("sign-up", authC.SignUp)
}

func userRoutes(r *gin.Engine) {
	userC := new(controllers.UserController)

	r.GET("/me", middlewares.Auth(), userC.Me)
}
