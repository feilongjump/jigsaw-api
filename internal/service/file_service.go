package service

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/dromara/carbon/v2"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

func (s *FileService) Save(file *multipart.FileHeader) (string, error) {
	fileType := resolveFileType(file)
	now := carbon.Now("Asia/Shanghai")
	date := now.ToDateString()
	timeText := strings.ReplaceAll(now.ToTimeString(), ":", "")
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		ext = ".bin"
	}

	randomText := randomSuffix()
	relativePath := filepath.Join(fileType, date, timeText+"_"+randomText+ext)
	fullPath := filepath.Join("static", relativePath)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return "/static/" + filepath.ToSlash(relativePath), nil
}

func resolveFileType(file *multipart.FileHeader) string {
	contentType := file.Header.Get("Content-Type")
	if contentType != "" {
		parts := strings.Split(contentType, "/")
		if len(parts) > 0 && parts[0] != "" {
			return parts[0]
		}
	}
	return "file"
}

func randomSuffix() string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return "00000000"
	}
	return hex.EncodeToString(buf)
}
