package router

import (
	"github.com/feilongjump/jigsaw-api/api/handler"
	"github.com/feilongjump/jigsaw-api/application/note"
	"github.com/feilongjump/jigsaw-api/infrastructure/repo_impl"
	"github.com/gin-gonic/gin"
)

func RegisterNoteRouter(r *gin.Engine) {
	noteRepo := repo_impl.NewNoteRepository()
	noteService := note.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHandler(noteService)

	group := r.Group("/notes")
	{
		group.POST("", noteHandler.Create)
		group.GET("", noteHandler.Index)
		group.GET("/:id", noteHandler.Show)
		group.PUT("/:id", noteHandler.Update)
		group.DELETE("/:id", noteHandler.Delete)
	}
}
