package handler

import (
	"github.com/feilongjump/jigsaw-api/application/user"
	"github.com/feilongjump/jigsaw-api/application/user/dto"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/gin_util"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *user.Service
}

func NewUserHandler(userService *user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register 用户注册
func (u *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !gin_util.BindJSON(c, &req) {
		return
	}

	if err := u.userService.Register(&req); err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, nil)
}

// UpdateAvatar 更新头像
func (u *UserHandler) UpdateAvatar(c *gin.Context) {
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		response.Fail(c, err_code.MalformedRequest)
		return
	}

	userID := c.GetUint64("user_id")
	resp, err := u.userService.UpdateAvatar(userID, fileHeader)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, resp)
}

// Login 用户登录
func (u *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !gin_util.BindJSON(c, &req) {
		return
	}

	resp, err := u.userService.Login(&req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, resp)
}

// GetProfile 获取当前用户信息
func (u *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint64("user_id")
	resp, err := u.userService.GetProfile(userID)
	if err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, resp)
}

// ChangePassword 修改密码
func (u *UserHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")
	if err := u.userService.ChangePassword(userID, &req); err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, nil)
}
