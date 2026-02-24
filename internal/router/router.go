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
	postHandler := handler.NewPostHandler()
	tagHandler := handler.NewTagHandler()

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
			users.PUT("/info", userHandler.UpdateUserInfo)
			users.PUT("/password", userHandler.ChangePassword)
		}

		files := r.Group("/files")
		files.Use(middleware.JWTAuth())
		{
			files.POST("/upload", fileHandler.Upload)
		}

		posts := r.Group("/posts")
		posts.Use(middleware.JWTAuth())
		{
			posts.GET("", postHandler.ListPosts)
			posts.POST("", postHandler.CreatePost)
			posts.GET("/:id", postHandler.GetPost)
			posts.PUT("/:id", postHandler.UpdatePost)
			posts.DELETE("/:id", postHandler.DeletePost)
		}

		tags := r.Group("/tags")
		tags.Use(middleware.JWTAuth())
		{
			tags.GET("", tagHandler.ListTags)
			tags.POST("", tagHandler.CreateTag)
			tags.PUT("/:id", tagHandler.UpdateTag)
			tags.DELETE("/:id", tagHandler.DeleteTag)
		}
	}

	return r
}
