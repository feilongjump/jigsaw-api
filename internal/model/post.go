package model

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type Post struct {
	ID              uint             `gorm:"primarykey" json:"id"`
	UserID          uint             `gorm:"index;not null;comment:用户ID" json:"user_id"`
	Title           string           `gorm:"not null;type:varchar(200);comment:标题" json:"title"`
	ContentMarkdown string           `gorm:"type:text;not null;comment:内容Markdown" json:"content_markdown"`
	ContentHTML     string           `gorm:"type:text;comment:内容HTML" json:"content_html"`
	Summary         string           `gorm:"type:text;comment:简介" json:"summary"`
	Cover           string           `gorm:"type:text;comment:封面" json:"cover"`
	Tags            []Tag            `gorm:"-" json:"tags"`
	CreatedAt       *carbon.DateTime `gorm:"autoCreateTime;comment:创建时间" json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt       *carbon.DateTime `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at" swaggertype:"string" format:"date-time"`
	DeletedAt       gorm.DeletedAt   `gorm:"index;comment:删除时间" json:"-"`
}
