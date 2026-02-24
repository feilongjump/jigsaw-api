package model

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type Tag struct {
	ID        uint             `gorm:"primarykey" json:"id"`
	Name      string           `gorm:"uniqueIndex;not null;type:varchar(50);comment:标签名" json:"name"`
	CreatedAt *carbon.DateTime `gorm:"autoCreateTime;comment:创建时间" json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt *carbon.DateTime `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at" swaggertype:"string" format:"date-time"`
	DeletedAt gorm.DeletedAt   `gorm:"index;comment:删除时间" json:"-"`
}
