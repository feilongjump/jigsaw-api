package repo

import "github.com/feilongjump/jigsaw-api/domain/entity"

// FileRepository 文件仓储接口
type FileRepository interface {
	// Create 创建文件记录
	Create(file *entity.File) error

	// GetFile 根据 ID 获取文件
	GetFile(id uint64) (*entity.File, error)

	// GetFileByPath 根据路径获取文件
	GetFileByPath(path string) (*entity.File, error)

	// FindFiles 根据 Owner 查找文件列表
	FindFiles(ownerType string, ownerID uint64) ([]*entity.File, error)

	// Delete 删除文件
	Delete(id uint64) error
}
