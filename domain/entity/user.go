package entity

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type User struct {
	ID uint64 `gorm:"primarykey" json:"id"`

	Username string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	CreatedAt *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}
