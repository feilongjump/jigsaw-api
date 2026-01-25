package router

import (
	"github.com/feilongjump/jigsaw-api/api/handler"
	"github.com/feilongjump/jigsaw-api/api/middleware"
	"github.com/feilongjump/jigsaw-api/application/ledger_category"
	"github.com/feilongjump/jigsaw-api/infrastructure/repo_impl"
	"github.com/gin-gonic/gin"
)

func RegisterLedgerCategoryRouter(r *gin.Engine) {
	repo := repo_impl.NewLedgerCategoryRepo()
	service := ledger_category.NewService(repo)
	h := handler.NewLedgerCategoryHandler(service)

	g := r.Group("/ledger/categories")
	g.Use(middleware.JWTAuth())
	{
		g.POST("", h.Create)
		g.GET("", h.Index)
		g.PUT("/:id", h.Update)
		g.DELETE("/:id", h.Delete)
	}
}
