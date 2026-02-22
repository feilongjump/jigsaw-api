package handler

import (
	"net/http"

	"jigsaw-api/internal/api/request"
	"jigsaw-api/internal/model"
	"jigsaw-api/internal/service"
	"jigsaw-api/pkg/response"
	"jigsaw-api/pkg/validator"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

// Register godoc
// @Summary 用户注册
// @Description 用户注册接口
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "注册信息"
// @Success 200 {object} response.Response "{"code": 200, "msg": "成功", "data": null}"
// @Failure 422 {object} response.Response "{"code": 422, "msg": "参数校验失败", "data": null, "errors": {...}}"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := h.authService.Register(user); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}

// Login godoc
// @Summary 用户登录
// @Description 用户登录接口
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "登录信息"
// @Success 200 {object} response.Response{data=string} "{"code": 200, "msg": "成功", "data": "token..."}"
// @Failure 422 {object} response.Response "{"code": 422, "msg": "参数校验失败", "data": null, "errors": {...}}"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	token, err := h.authService.Login(req.Account, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, gin.H{"token": token})
}
