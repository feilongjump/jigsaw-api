package router

import (
	"github.com/feilongjump/jigsaw-api/api/handler"
	"github.com/feilongjump/jigsaw-api/api/middleware"
	"github.com/feilongjump/jigsaw-api/application/ledger_record"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/infrastructure/repo_impl"
	"github.com/gin-gonic/gin"
)

func RegisterLedgerRecordRouter(r *gin.Engine) {
	recordRepo := repo_impl.NewLedgerRecordRepo()
	walletRepo := repo_impl.NewUserWalletRepo()
	tm := db.NewTransactionManager()
	service := ledger_record.NewService(recordRepo, walletRepo, tm)
	h := handler.NewLedgerRecordHandler(service)

	g := r.Group("/ledger/records")
	g.Use(middleware.JWTAuth())
	{
		g.POST("", h.Create)
		g.GET("", h.Index)
		g.PUT("/:id", h.Update)
		g.DELETE("/:id", h.Delete)
	}
}
