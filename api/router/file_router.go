package router

import (
	"github.com/feilongjump/jigsaw-api/api/handler"
	"github.com/feilongjump/jigsaw-api/api/middleware"
	"github.com/feilongjump/jigsaw-api/application/file"
	"github.com/feilongjump/jigsaw-api/infrastructure/repo_impl"
	"github.com/gin-gonic/gin"
)

func RegisterFileRouter(r *gin.Engine) {
	fileRepo := repo_impl.NewFileRepository()
	fileService := file.NewFileService(fileRepo)
	fileHandler := handler.NewFileHandler(fileService)

	group := r.Group("/files")
	group.Use(middleware.JWTAuth())
	{
		group.POST("/upload", fileHandler.Upload)
		group.POST("/delete", fileHandler.Delete)
	}
}
