package handler

import (
	"github.com/feilongjump/jigsaw-api/application/user_wallet"
	"github.com/feilongjump/jigsaw-api/application/user_wallet/dto"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/gin_util"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserWalletHandler struct {
	service *user_wallet.Service
}

func NewUserWalletHandler(service *user_wallet.Service) *UserWalletHandler {
	return &UserWalletHandler{service: service}
}

func (h *UserWalletHandler) Create(c *gin.Context) {
	var req dto.CreateUserWalletReq
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	createdWallet, err := h.service.Create(userID, req)
	if err != nil {
		response.Fail(c, err_code.UserWalletCreateFailed)
		return
	}

	response.Success(c, createdWallet)
}

func (h *UserWalletHandler) Index(c *gin.Context) {
	userID := c.GetUint64("user_id")
	walletsResp, err := h.service.FindUserWallets(userID)
	if err != nil {
		response.Fail(c, err_code.UserWalletGetFailed)
		return
	}

	response.Success(c, walletsResp)
}

func (h *UserWalletHandler) Delete(c *gin.Context) {
	var req dto.UserWalletURIReq
	if !gin_util.BindURI(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := h.service.Delete(userID, req.ID); err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, nil)
}

func (h *UserWalletHandler) Update(c *gin.Context) {
	var reqURI dto.UserWalletURIReq
	if !gin_util.BindURI(c, &reqURI) {
		return
	}

	var reqBody dto.UpdateUserWalletReq
	if !gin_util.BindJSON(c, &reqBody) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := h.service.Update(userID, reqURI.ID, reqBody); err != nil {
		response.Fail(c, err_code.UserWalletUpdateFailed)
		return
	}

	response.Success(c, nil)
}
