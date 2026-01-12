package entity

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type User struct {
	ID uint64 `gorm:"primarykey" json:"id"`

	Username string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	Files []File `gorm:"polymorphic:Owner;polymorphicValue:users" json:"files"`

	CreatedAt *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}

// BeforeDelete GORM 钩子，用于同步删除关联文件
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	var files []File
	// 查找与该 User 关联的文件
	if err := tx.Where("owner_type = ? AND owner_id = ?", "users", u.ID).Find(&files).Error; err != nil {
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
