package controllers

import (
	"github.com/feilongjump/jigsaw-api/app/http/responses"
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (h *HomeController) Index(ctx *gin.Context) {

	responses.Success(ctx, nil)
}
