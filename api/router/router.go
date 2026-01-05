package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {

	r := gin.Default()

	RegisterNoteRouter(r)
	RegisterUserRouter(r)

	return r
}
