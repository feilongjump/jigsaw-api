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

	CreatedAt *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}
