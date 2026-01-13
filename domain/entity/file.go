package entity

import (
	"os"

	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

// FileType 定义文件类型
type FileType string

const (
	FileTypeImage    FileType = "image"
	FileTypeVideo    FileType = "video"
	FileTypeDocument FileType = "document"
	TypeText         FileType = "text"
	FileTypeOther    FileType = "other"
)

// File 文件实体
type File struct {
	ID uint64 `gorm:"primarykey" json:"id"`

	Name     string   `gorm:"type:varchar(255);not null" json:"name"`
	Path     string   `gorm:"type:varchar(255);not null;unique" json:"path"`
	Type     FileType `gorm:"type:varchar(50);not null" json:"type"`
	Size     int64    `gorm:"not null" json:"size"`
	MimeType string   `gorm:"type:varchar(100)" json:"mime_type"`

	// OwnerType 标识该文件属于哪个业务模块，例如 "notes", "users"
	OwnerType string `gorm:"type:varchar(50);index" json:"owner_type"`
	// OwnerID 标识该文件所属业务模块的 ID，例如 note_id, user_id
	OwnerID uint64 `gorm:"index" json:"owner_id"`

	// UploaderID 上传该文件的用户 ID
	UploaderID uint64 `gorm:"index" json:"uploader_id"`

	CreatedAt *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}

// BeforeDelete GORM 钩子，用于在删除数据库记录前删除物理文件
func (f *File) BeforeDelete(tx *gorm.DB) (err error) {
	// 删除磁盘上的文件
	// 忽略文件不存在的错误
	_ = os.Remove(f.Path)

	return nil
}
