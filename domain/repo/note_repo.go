package repo

import "github.com/feilongjump/jigsaw-api/domain/entity"

// NoteRepository 接口（定义数据访问规则）
type NoteRepository interface {
	// Create 创建
	Create(note *entity.Note) error

	// GetNote 根据 ID 查询
	GetNote(id, userID uint64) (*entity.Note, error)

	// FindNotes 查询列表
	FindNotes(page, size int, userID uint64) ([]*entity.Note, int64, error)

	// Update 更新
	Update(id, userID uint64, note *entity.Note) error

	// Delete 删除
	Delete(id, userID uint64) (error, int64)
}
