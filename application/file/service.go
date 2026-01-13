package file

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/feilongjump/jigsaw-api/application/file/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
)

type Service struct {
	fileRepo repo.FileRepository
}

func NewFileService(fileRepo repo.FileRepository) *Service {
	return &Service{
		fileRepo: fileRepo,
	}
}

func (s *Service) Upload(fileHeader *multipart.FileHeader, ownerType string, ownerID uint64, uploaderID uint64) (*dto.FileResponse, error) {
	// 1. 打开文件
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// 2. 确定文件类型
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	fileType := getFileType(ext)

	// 3. 生成路径: static/type/date/
	dateStr := time.Now().Format("2006-01-02")
	uploadDir := filepath.Join("static", string(fileType), dateStr)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, err
	}

	// 4. 生成文件名: H-i-s_随机数.ext
	now := time.Now()
	timestampStr := now.Format("15-04-05")
	// 使用本地随机源，避免全局锁竞争，并确保每次运行都有不同的随机数序列
	rng := rand.New(rand.NewSource(now.UnixNano()))
	randomNum := rng.Intn(90000) + 10000 // 生成 10000-99999 之间的 5 位随机数
	newFilename := fmt.Sprintf("%s_%d%s", timestampStr, randomNum, ext)
	dstPath := filepath.Join(uploadDir, newFilename)

	// 5. 保存文件
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	// 6. 保存到数据库
	fileEntity := &entity.File{
		Name:       fileHeader.Filename,
		Path:       dstPath,
		Type:       fileType,
		Size:       fileHeader.Size,
		MimeType:   fileHeader.Header.Get("Content-Type"),
		OwnerType:  ownerType,
		OwnerID:    ownerID,
		UploaderID: uploaderID,
	}

	if err := s.fileRepo.Create(fileEntity); err != nil {
		// 如果数据库保存失败，清理文件
		os.Remove(dstPath)
		return nil, err
	}

	// 生成访问 URL (去掉 static 前缀，并确保使用 /)
	urlPath := strings.ReplaceAll(dstPath, "\\", "/")
	urlPath = strings.TrimPrefix(urlPath, "static")
	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}

	return &dto.FileResponse{
		ID:        fileEntity.ID,
		Name:      fileEntity.Name,
		Url:       urlPath,
		Size:      fileEntity.Size,
		MimeType:  fileEntity.MimeType,
		OwnerType: fileEntity.OwnerType,
		OwnerID:   fileEntity.OwnerID,
		CreatedAt: fileEntity.CreatedAt,
	}, nil
}

// Delete 删除文件
func (s *Service) Delete(path string, userID uint64, ownerType string, ownerID uint64) error {
	// 1. 还原存储路径 (static + url path)
	// path 来自 URL，如 "/image/2023/xxx.jpg"，需要拼接 static 才能找到文件
	storagePath := filepath.Join("static", path)

	// 1. 查询文件
	fileEntity, err := s.fileRepo.GetFileByPath(storagePath)
	if err != nil {
		return err
	}

	// 2. 权限校验: 只能删除自己上传的文件
	if fileEntity.UploaderID != userID {
		return err_code.FileDeleteForbidden
	}

	// 3. 模块校验: 确保删除的是指定模块下的文件 (双重确认)
	if fileEntity.OwnerType != ownerType || fileEntity.OwnerID != ownerID {
		return err_code.ValidationFailed
	}

	// 4. 删除数据库记录 (触发 BeforeDelete Hook 删除物理文件)
	return s.fileRepo.Delete(fileEntity.ID)
}

func getFileType(ext string) entity.FileType {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
		return entity.FileTypeImage
	case ".mp4", ".avi", ".mov", ".wmv":
		return entity.FileTypeVideo
	case ".doc", ".docx", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx":
		return entity.FileTypeDocument
	case ".txt", ".md", ".json", ".xml":
		return entity.TypeText
	default:
		return entity.FileTypeOther
	}
}
