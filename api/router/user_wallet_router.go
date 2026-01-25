package router

import (
	"github.com/feilongjump/jigsaw-api/api/handler"
	"github.com/feilongjump/jigsaw-api/api/middleware"
	"github.com/feilongjump/jigsaw-api/application/user_wallet"
	"github.com/feilongjump/jigsaw-api/infrastructure/repo_impl"
	"github.com/gin-gonic/gin"
)

func RegisterUserWalletRouter(r *gin.Engine) {
	repo := repo_impl.NewUserWalletRepo()
	recordRepo := repo_impl.NewLedgerRecordRepo()
	service := user_wallet.NewService(repo, recordRepo)
	h := handler.NewUserWalletHandler(service)

	g := r.Group("/users/wallets")
	g.Use(middleware.JWTAuth())
	{
		g.POST("", h.Create)
		g.GET("", h.Index)
		g.PUT("/:id", h.Update)
		g.DELETE("/:id", h.Delete)
	}
}
