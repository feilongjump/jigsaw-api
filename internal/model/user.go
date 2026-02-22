package model

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID          uint             `gorm:"primarykey" json:"id"`
	Username    string           `gorm:"uniqueIndex;not null;comment:用户名" json:"username" binding:"required,min=3,max=20" label:"用户名"`
	Password    string           `gorm:"not null;comment:密码" json:"-" binding:"required,min=6" label:"密码"`
	Email       string           `gorm:"uniqueIndex;not null;comment:邮箱" json:"email" binding:"required,email" label:"邮箱"`
	Phone       string           `gorm:"uniqueIndex;default:null;comment:手机号" json:"phone" label:"手机号"`
	Avatar      string           `gorm:"comment:头像" json:"avatar" label:"头像"`
	LastLoginAt *carbon.DateTime `gorm:"comment:最后登录时间" json:"last_login_at" swaggertype:"string" format:"date-time" label:"最后登录时间"`
	CreatedAt   *carbon.DateTime `gorm:"autoCreateTime;comment:创建时间" json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt   *carbon.DateTime `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at" swaggertype:"string" format:"date-time"`
	DeletedAt   gorm.DeletedAt   `gorm:"index;comment:删除时间" json:"-"`
}
