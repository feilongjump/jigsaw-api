package handler

import (
	"github.com/feilongjump/jigsaw-api/application/ledger_record"
	"github.com/feilongjump/jigsaw-api/application/ledger_record/dto"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/gin_util"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type LedgerRecordHandler struct {
	service *ledger_record.Service
}

func NewLedgerRecordHandler(service *ledger_record.Service) *LedgerRecordHandler {
	return &LedgerRecordHandler{service: service}
}

func (h *LedgerRecordHandler) Create(c *gin.Context) {
	var req dto.CreateLedgerRecordReq
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	created, err := h.service.Create(c.Request.Context(), userID, req)
	if err != nil {
		response.Fail(c, err_code.LedgerRecordCreateFailed)
		return
	}

	response.Success(c, created)
}

func (h *LedgerRecordHandler) Update(c *gin.Context) {
	var reqURI dto.LedgerRecordURIReq
	if !gin_util.BindURI(c, &reqURI) {
		return
	}

	var req dto.UpdateLedgerRecordReq
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	updated, err := h.service.Update(c.Request.Context(), userID, reqURI.ID, req)
	if err != nil {
		response.Fail(c, err_code.LedgerRecordUpdateFailed)
		return
	}

	response.Success(c, updated)
}

func (h *LedgerRecordHandler) Delete(c *gin.Context) {
	var req dto.LedgerRecordURIReq
	if !gin_util.BindURI(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := h.service.Delete(c.Request.Context(), userID, req.ID); err != nil {
		response.Fail(c, err_code.LedgerRecordDeleteFailed)
		return
	}

	response.Success(c, nil)
}

func (h *LedgerRecordHandler) Index(c *gin.Context) {
	var req dto.ListLedgerRecordReq
	if !gin_util.BindQuery(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	resp, err := h.service.FindLedgerRecords(userID, req)
	if err != nil {
		response.Fail(c, err_code.LedgerRecordGetFailed)
		return
	}

	response.Success(c, resp)
}
