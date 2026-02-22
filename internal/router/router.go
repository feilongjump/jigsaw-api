package router

import (
	"jigsaw-api/internal/handler"
	"jigsaw-api/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.CORS())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.Recovery())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/static", "./static")

	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()
	fileHandler := handler.NewFileHandler()

	{
		auth := r.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		users := r.Group("/users")
		users.Use(middleware.JWTAuth())
		{
			r.GET("/me", middleware.JWTAuth(), userHandler.GetMe)
			users.GET("", userHandler.FindUsers)
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
			users.PUT("/info", userHandler.UpdateUserInfo)
		}

		files := r.Group("/files")
		files.Use(middleware.JWTAuth())
		{
			files.POST("/upload", fileHandler.Upload)
		}
	}

	return r
}
