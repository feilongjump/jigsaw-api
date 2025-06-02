package user

import (
	"github.com/feilongjump/jigsaw-api/app/models"
	"github.com/feilongjump/jigsaw-api/plugins/database"
	"github.com/feilongjump/jigsaw-api/utils"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(16);not null;index" json:"name,omitempty"`
	Email    string `gorm:"type:varchar(128);default:null;uniqueIndex" json:"email,omitempty"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar,omitempty"`

	models.CommonTimestampsField
}

func (u *User) Create() error {
	return database.DB.Create(&u).Error
}

func (u *User) ComparePassword(password string) bool {
	return utils.BcryptCheck(u.Password, password)
}
