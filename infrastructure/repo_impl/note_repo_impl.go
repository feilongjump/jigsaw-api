package repo_impl

import (
	"errors"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"gorm.io/gorm"
)

// noteRepositoryImpl Note 仓储实现（MySQL）
type noteRepositoryImpl struct {
	db *gorm.DB
}

// NewNoteRepository 创建 Note 仓储实例
func NewNoteRepository() repo.NoteRepository {
	return &noteRepositoryImpl{
		db: db.Get(),
	}
}

// Create 创建 Note
func (r *noteRepositoryImpl) Create(note *entity.Note) error {
	return r.db.Create(note).Error
}

// GetNote 根据 ID 查询 Note
func (r *noteRepositoryImpl) GetNote(id, userID uint64) (*entity.Note, error) {
	var note entity.Note
	err := r.db.
		Where("user_id", userID).
		First(&note, id).
		Error
	if err != nil {
		// 当数据不存在时，将返回自定义的数据不存在错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.NoteNotFound
		}
		return nil, err
	}
	return &note, nil
}

// FindNotes 查询 Note 列表
func (r *noteRepositoryImpl) FindNotes(page, size int, userID uint64) ([]*entity.Note, int64, error) {
	var notes []*entity.Note
	var total int64
	err := r.db.
		Where("user_id", userID).
		Model(&entity.Note{}).
		Count(&total).
		Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.
		Where("user_id", userID).
		Offset((page - 1) * size).
		Limit(size).
		Order("created_at desc").
		Find(&notes).
		Error
	if err != nil {
		return nil, 0, err
	}
	return notes, total, nil
}

// Update 更新 Note
func (r *noteRepositoryImpl) Update(id, userID uint64, note *entity.Note) error {
	return r.db.
		Where("user_id", userID).
		Model(&entity.Note{
			ID: id,
		}).
		Updates(note).
		Error
}

// Delete 删除 Note
func (r *noteRepositoryImpl) Delete(id, userID uint64) (error, int64) {
	result := r.db.
		Where("user_id", userID).
		Delete(&entity.Note{
			ID: id,
		})

	return result.Error, result.RowsAffected
}
