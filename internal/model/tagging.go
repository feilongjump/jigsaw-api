package model

import "github.com/dromara/carbon/v2"

type Tagging struct {
	ID           uint             `gorm:"primarykey" json:"id"`
	TagID        uint             `gorm:"index;not null;comment:标签ID" json:"tag_id"`
	ResourceType string           `gorm:"index;not null;type:varchar(50);comment:资源类型" json:"resource_type"`
	ResourceID   uint             `gorm:"index;not null;comment:资源ID" json:"resource_id"`
	CreatedAt    *carbon.DateTime `gorm:"autoCreateTime;comment:创建时间" json:"created_at" swaggertype:"string" format:"date-time"`
}
