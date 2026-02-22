package handler

import (
	"net/http"

	"jigsaw-api/internal/api/request"
	"jigsaw-api/internal/service"
	"jigsaw-api/pkg/response"
	"jigsaw-api/pkg/validator"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService *service.FileService
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		fileService: service.NewFileService(),
	}
}

// Upload godoc
// @Summary 上传文件
// @Description 上传文件接口
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} response.Response{data=response.UploadFileResponse} "{"code": 200, "msg": "成功", "data": {"path": "/static/image/2026-02-21/120102_ab12cd34.png"}}"
// @Failure 422 {object} response.Response "{"code": 422, "msg": "参数校验失败", "data": null, "errors": {...}}"
// @Router /files/upload [post]
// @Security ApiKeyAuth
func (h *FileHandler) Upload(c *gin.Context) {
	var req request.UploadFileRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ValidationError(c, validator.Translate(err, &req))
		return
	}

	path, err := h.fileService.Save(req.File)
	if err != nil {
		response.ErrorWithStatus(c, http.StatusInternalServerError, 500, "文件上传失败")
		return
	}

	response.Success(c, response.UploadFileResponse{Path: path})
}
