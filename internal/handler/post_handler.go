package handler

import (
	"net/http"
	"strconv"

	"jigsaw-api/internal/api/request"
	"jigsaw-api/internal/model"
	"jigsaw-api/internal/service"
	"jigsaw-api/pkg/response"
	"jigsaw-api/pkg/validator"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		postService: service.NewPostService(),
	}
}

// ListPosts godoc
// @Summary 获取文章列表
// @Description 获取当前登录用户文章列表
// @Tags Post
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} response.Response{data=[]model.Post} "{"code": 200, "msg": "成功", "data": [...]}"
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /posts [get]
// @Security ApiKeyAuth
func (h *PostHandler) ListPosts(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

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

	posts, total, err := h.postService.FindByUser(userID.(uint), page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "获取文章列表失败")
		return
	}

	response.Success(c, gin.H{
		"list":  posts,
		"total": total,
	})
}

// CreatePost godoc
// @Summary 创建文章
// @Description 创建文章
// @Tags Post
// @Accept json
// @Produce json
// @Param request body request.CreatePostRequest true "文章信息"
// @Success 200 {object} response.Response{data=model.Post}
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response "{\"code\": 422, \"msg\": \"参数校验失败\", \"data\": null, \"errors\": {...}}"
// @Router /posts [post]
// @Security ApiKeyAuth
func (h *PostHandler) CreatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	var req request.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	post := &model.Post{
		UserID:          userID.(uint),
		Title:           req.Title,
		ContentMarkdown: req.ContentMarkdown,
		ContentHTML:     req.ContentHTML,
		Summary:         req.Summary,
		Cover:           req.Cover,
	}

	created, err := h.postService.Create(post, req.Tags)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "创建文章失败")
		return
	}

	response.Success(c, created)
}

// GetPost godoc
// @Summary 获取文章详情
// @Description 根据ID获取文章详情
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} response.Response{data=model.Post}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /posts/{id} [get]
// @Security ApiKeyAuth
func (h *PostHandler) GetPost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "文章ID无效")
		return
	}

	post, err := h.postService.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "文章不存在")
		return
	}

	if post.UserID != userID.(uint) {
		response.Error(c, http.StatusForbidden, "无权限访问该文章")
		return
	}

	response.Success(c, post)
}

// UpdatePost godoc
// @Summary 更新文章
// @Description 更新文章内容
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Param request body request.UpdatePostRequest true "更新信息"
// @Success 200 {object} response.Response{data=model.Post}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 422 {object} response.Response "{\"code\": 422, \"msg\": \"参数校验失败\", \"data\": null, \"errors\": {...}}"
// @Router /posts/{id} [put]
// @Security ApiKeyAuth
func (h *PostHandler) UpdatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "文章ID无效")
		return
	}

	post, err := h.postService.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "文章不存在")
		return
	}

	if post.UserID != userID.(uint) {
		response.Error(c, http.StatusForbidden, "无权限访问该文章")
		return
	}

	var req request.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.ContentMarkdown != nil {
		post.ContentMarkdown = *req.ContentMarkdown
	}
	if req.ContentHTML != nil {
		post.ContentHTML = *req.ContentHTML
	}
	if req.Summary != nil {
		post.Summary = *req.Summary
	}
	if req.Cover != nil {
		post.Cover = *req.Cover
	}

	updated, err := h.postService.Update(post, req.Tags)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "更新文章失败")
		return
	}

	response.Success(c, updated)
}

// DeletePost godoc
// @Summary 删除文章
// @Description 删除文章
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /posts/{id} [delete]
// @Security ApiKeyAuth
func (h *PostHandler) DeletePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "用户身份信息不存在")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "文章ID无效")
		return
	}

	post, err := h.postService.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "文章不存在")
		return
	}

	if post.UserID != userID.(uint) {
		response.Error(c, http.StatusForbidden, "无权限访问该文章")
		return
	}

	if err := h.postService.Delete(post.ID); err != nil {
		response.Error(c, http.StatusInternalServerError, "删除文章失败")
		return
	}

	response.Success(c, nil)
}
