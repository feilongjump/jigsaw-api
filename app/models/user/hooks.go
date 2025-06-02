package user

import (
	"errors"
	"github.com/feilongjump/jigsaw-api/utils"
	"gorm.io/gorm"
)

// BeforeSave 密码加密
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" {
		return errors.New("密码不能为空")
	}

	u.Password = utils.BcryptHash(u.Password)
	return nil
}
