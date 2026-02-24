package handler

import (
	"net/http"
	"strconv"

	"jigsaw-api/internal/api/request"
	"jigsaw-api/internal/service"
	"jigsaw-api/pkg/response"
	"jigsaw-api/pkg/validator"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler() *TagHandler {
	return &TagHandler{
		tagService: service.NewTagService(),
	}
}

// ListTags godoc
// @Summary 获取标签列表
// @Description 获取标签列表
// @Tags Tag
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} response.Response "{"code": 200, "msg": "成功", "data": [...]}"
// @Failure 400 {object} response.Response
// @Router /tags [get]
// @Security ApiKeyAuth
func (h *TagHandler) ListTags(c *gin.Context) {
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

	tags, total, err := h.tagService.List(page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "获取标签列表失败")
		return
	}

	response.Success(c, gin.H{
		"list":  tags,
		"total": total,
	})
}

// CreateTag godoc
// @Summary 创建标签
// @Description 创建标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param request body request.CreateTagRequest true "标签信息"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response "{\"code\": 422, \"msg\": \"参数校验失败\", \"data\": null, \"errors\": {...}}"
// @Router /tags [post]
// @Security ApiKeyAuth
func (h *TagHandler) CreateTag(c *gin.Context) {
	var req request.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	tag, err := h.tagService.Create(req.Name)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "创建标签失败")
		return
	}

	response.Success(c, tag)
}

// UpdateTag godoc
// @Summary 更新标签
// @Description 更新标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param id path int true "标签ID"
// @Param request body request.UpdateTagRequest true "标签信息"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 422 {object} response.Response "{\"code\": 422, \"msg\": \"参数校验失败\", \"data\": null, \"errors\": {...}}"
// @Router /tags/{id} [put]
// @Security ApiKeyAuth
func (h *TagHandler) UpdateTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "标签ID无效")
		return
	}

	var req request.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	tag, err := h.tagService.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "标签不存在")
		return
	}
	tag.Name = req.Name

	if err := h.tagService.Update(tag); err != nil {
		response.Error(c, http.StatusInternalServerError, "更新标签失败")
		return
	}

	response.Success(c, tag)
}

// DeleteTag godoc
// @Summary 删除标签
// @Description 删除标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /tags/{id} [delete]
// @Security ApiKeyAuth
func (h *TagHandler) DeleteTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "标签ID无效")
		return
	}

	if err := h.tagService.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, "删除标签失败")
		return
	}

	response.Success(c, nil)
}
