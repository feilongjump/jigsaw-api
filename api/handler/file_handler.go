package handler

import (
	"github.com/feilongjump/jigsaw-api/application/file"
	"github.com/feilongjump/jigsaw-api/application/file/dto"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/gin_util"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService *file.Service
}

func NewFileHandler(fileService *file.Service) *FileHandler {
	return &FileHandler{
		fileService: fileService,
	}
}

// Upload 处理文件上传
func (h *FileHandler) Upload(c *gin.Context) {
	// 获取上传的文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, err_code.MalformedRequest)
		return
	}

	// 绑定其他参数
	var req dto.UploadFileRequest
	if !gin_util.BindMultipart(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")

	// 处理 owner_id 逻辑
	var ownerID uint64
	if req.OwnerType == "users" {
		ownerID = userID
	} else {
		// 如果不是 user 模块，owner_id 必填
		if req.OwnerID == 0 {
			response.ValidateFail(c, map[string][]string{
				"owner_id": {"owner_id 不能为空"},
			})
			return
		}
		ownerID = req.OwnerID
	}

	resp, err := h.fileService.Upload(fileHeader, req.OwnerType, ownerID, userID)
	if err != nil {
		response.Fail(c, err_code.FileUploadFailed)
		return
	}

	response.Success(c, resp)
}

// Delete 删除文件
func (h *FileHandler) Delete(c *gin.Context) {
	var req dto.DeleteFileRequest
	// 绑定 JSON 参数
	if !gin_util.BindJSON(c, &req) {
		return
	}

	userID := c.GetUint64("user_id")

	// 处理 Path 开头的 / (虽然 JSON 传递通常不带，但为了兼容性处理一下)
	path := req.Path
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	if err := h.fileService.Delete(path, userID, req.OwnerType, req.OwnerID); err != nil {
		response.Fail(c, err)
		return
	}

	response.Success(c, nil)
}
