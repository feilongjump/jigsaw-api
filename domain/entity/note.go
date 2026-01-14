package entity

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

// Note 实体（领域模型）
type Note struct {
	ID     uint64 `gorm:"primarykey" json:"id"`
	UserID uint64 `gorm:"index;not null" json:"user_id"`

	Content string `gorm:"type:longtext"`

	Files []File `gorm:"polymorphic:Owner;polymorphicValue:notes" json:"files"`

	PinnedAt *carbon.DateTime `gorm:"type:datetime;index" json:"pinned_at"` // 置顶时间，为空表示未置顶

	CreatedAt *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}

// BeforeDelete GORM 钩子，用于同步删除关联文件
func (n *Note) BeforeDelete(tx *gorm.DB) (err error) {
	var files []File
	// 查找与该 Note 关联的文件
	if err := tx.Where("owner_type = ? AND owner_id = ?", "notes", n.ID).Find(&files).Error; err != nil {
		return err
	}

	for _, f := range files {
		// 从数据库删除文件记录
		// 这将触发 File.BeforeDelete 钩子，从而删除物理文件
		if err := tx.Delete(&f).Error; err != nil {
			return err
		}
	}
	return nil
}
