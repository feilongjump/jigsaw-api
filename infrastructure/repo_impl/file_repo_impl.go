package repo_impl

import (
	"errors"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"gorm.io/gorm"
)

type fileRepositoryImpl struct {
	db *gorm.DB
}

func NewFileRepository() repo.FileRepository {
	return &fileRepositoryImpl{
		db: db.Get(),
	}
}

func (r *fileRepositoryImpl) Create(file *entity.File) error {
	return r.db.Create(file).Error
}

// GetFile 根据 ID 获取文件
func (r *fileRepositoryImpl) GetFile(id uint64) (*entity.File, error) {
	var file entity.File
	err := r.db.First(&file, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.FileNotFound
		}
		return nil, err
	}
	return &file, nil
}

// GetFileByPath 根据路径获取文件
func (r *fileRepositoryImpl) GetFileByPath(path string) (*entity.File, error) {
	var file entity.File
	err := r.db.Where("path = ?", path).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.FileNotFound
		}
		return nil, err
	}
	return &file, nil
}

// FindFiles 根据 Owner 查找文件列表
func (r *fileRepositoryImpl) FindFiles(ownerType string, ownerID uint64) ([]*entity.File, error) {
	var files []*entity.File
	err := r.db.Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).Find(&files).Error
	return files, err
}

// BindFiles 绑定文件到 Owner
func (r *fileRepositoryImpl) BindFiles(fileIDs []uint64, userID uint64, ownerType string, ownerID uint64) error {
	if len(fileIDs) == 0 {
		return nil
	}
	// 只更新属于该用户的文件，防止恶意篡改他人文件
	return r.db.Model(&entity.File{}).
		Where("id IN ? AND uploader_id = ?", fileIDs, userID).
		Updates(map[string]interface{}{
			"owner_type": ownerType,
			"owner_id":   ownerID,
		}).Error
}

// Delete 删除文件
func (r *fileRepositoryImpl) Delete(id uint64) error {
	// 先查询出文件信息，确保 BeforeDelete 钩子能获取到 Path
	var file entity.File
	if err := r.db.First(&file, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // 如果文件已经不存在，视为删除成功
		}
		return err
	}
	return r.db.Delete(&file).Error
}
