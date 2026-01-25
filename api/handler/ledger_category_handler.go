package handler

import (
	"github.com/feilongjump/jigsaw-api/application/ledger_category"
	"github.com/feilongjump/jigsaw-api/application/ledger_category/dto"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/gin_util"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type LedgerCategoryHandler struct {
	service *ledger_category.Service
}

func NewLedgerCategoryHandler(service *ledger_category.Service) *LedgerCategoryHandler {
	return &LedgerCategoryHandler{service: service}
}

func (h *LedgerCategoryHandler) Create(c *gin.Context) {
	var req dto.CreateLedgerCategoryReq
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := h.service.Create(userID, req); err != nil {
		response.Fail(c, err_code.LedgerCategoryCreateFailed)
		return
	}

	response.Success(c, nil)
}

func (h *LedgerCategoryHandler) Index(c *gin.Context) {
	userID := c.GetUint64("user_id")
	categories, err := h.service.FindLedgerCategories(userID)
	if err != nil {
		response.Fail(c, err_code.LedgerCategoryGetFailed)
		return
	}

	response.Success(c, categories)
}

func (h *LedgerCategoryHandler) Delete(c *gin.Context) {
	var req dto.LedgerCategoryURIReq
	if !gin_util.BindURI(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := h.service.Delete(userID, req.ID); err != nil {
		response.Fail(c, err_code.LedgerCategoryDeleteFailed)
		return
	}

	response.Success(c, nil)
}

func (h *LedgerCategoryHandler) Update(c *gin.Context) {
	var reqURI dto.LedgerCategoryURIReq
	if !gin_util.BindURI(c, &reqURI) {
		return
	}

	var req dto.UpdateLedgerCategoryReq
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := h.service.Update(userID, reqURI.ID, req); err != nil {
		response.Fail(c, err_code.LedgerCategoryUpdateFailed)
		return
	}

	updatedCategory, err := h.service.GetLedgerCategory(userID, reqURI.ID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, updatedCategory)
}
