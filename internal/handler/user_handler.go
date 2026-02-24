package handler

import (
	"net/http"
	"strconv"

	"jigsaw-api/internal/api/request"
	"jigsaw-api/internal/model"
	"jigsaw-api/internal/service"
	"jigsaw-api/pkg/response"
	"jigsaw-api/pkg/validator"

	// imported for swagger

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// FindUsers godoc
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags User
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} response.Response{data=[]model.User} "{"code": 200, "msg": "成功", "data": [...]}"
// @Failure 400 {object} response.Response
// @Router /users [get]
// @Security ApiKeyAuth
func (h *UserHandler) FindUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	users, total, err := h.userService.FindUsers(page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "获取用户列表失败")
		return
	}

	response.Success(c, gin.H{
		"list":  users,
		"total": total,
	})
}

// GetUser godoc
// @Summary 获取单个用户
// @Description 根据ID获取用户信息
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=model.User}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "用户ID无效")
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	response.Success(c, user)
}

// GetMe godoc
// @Summary 获取个人信息
// @Description 获取当前登录用户信息
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=model.User}
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /me [get]
// @Security ApiKeyAuth
func (h *UserHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	user, err := h.userService.GetUser(userID.(uint))
	if err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	response.Success(c, user)
}

// CreateUser godoc
// @Summary 创建用户
// @Description 创建新用户
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "用户信息"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response "{\"code\": 422, \"msg\": \"参数校验失败\", \"data\": null, \"errors\": {...}}"
// @Router /users [post]
// @Security ApiKeyAuth
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := h.userService.CreateUser(user); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}

// UpdateUser godoc
// @Summary 更新用户
// @Description 更新用户信息
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param request body request.UpdateUserRequest true "更新信息"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response "{\"code\": 422, \"msg\": \"参数校验失败\", \"data\": null, \"errors\": {...}}"
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	updates := make(map[string]interface{})
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}

	if err := h.userService.UpdateUser(uint(id), updates); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}

// UpdateUserInfo godoc
// @Summary 更新个人信息
// @Description 更新当前登录用户的信息 (头像, 邮箱等)
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.UpdateUserRequest true "更新信息"
// @Success 200 {object} response.Response "{"code": 200, "msg": "成功", "data": null}"
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response "{"code": 422, "msg": "参数校验失败", "data": null, "errors": {...}}"
// @Router /users/info [put]
// @Security ApiKeyAuth
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	updates := make(map[string]interface{})
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}

	if err := h.userService.UpdateUser(userID.(uint), updates); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}

// ChangePassword godoc
// @Summary 修改密码
// @Description 登录后修改密码
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.ChangePasswordRequest true "修改密码信息"
// @Success 200 {object} response.Response "{"code": 200, "msg": "成功", "data": null}"
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response "{"code": 422, "msg": "参数校验失败", "data": null, "errors": {...}}"
// @Router /users/password [put]
// @Security ApiKeyAuth
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	var req request.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	if err := h.userService.ChangePassword(userID.(uint), req.OldPassword, req.NewPassword); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}
