package router

import (
	"github.com/feilongjump/jigsaw-api/api/handler"
	"github.com/feilongjump/jigsaw-api/api/middleware"
	"github.com/feilongjump/jigsaw-api/application/user"
	"github.com/feilongjump/jigsaw-api/infrastructure/repo_impl"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.Engine) {
	// 依赖注入
	userRepo := repo_impl.NewUserRepository()
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 路由分组
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
	}

	// 需要认证的路由
	userGroup := r.Group("/user")
	userGroup.Use(middleware.JWTAuth())
	{
		userGroup.POST("/change-password", userHandler.ChangePassword)
	}
}
