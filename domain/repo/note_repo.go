package repo

import "github.com/feilongjump/jigsaw-api/domain/entity"

// NoteRepository 接口（定义数据访问规则）
type NoteRepository interface {
	// Create 创建
	Create(note *entity.Note) error

	// GetNote 根据 ID 查询
	GetNote(id uint64) (*entity.Note, error)

	// FindNotes 查询列表
	FindNotes(page, size int) ([]*entity.Note, int64, error)

	// Update 更新
	Update(id uint64, note *entity.Note) error

	// Delete 删除
	Delete(id uint64) error
}
