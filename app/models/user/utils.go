package user

import "github.com/feilongjump/jigsaw-api/plugins/database"

// FindByMulti 根据用户名称或邮箱查找用户
func FindByMulti(str string) (user User, err error) {

	err = database.DB.Where("name", str).
		Or("email", str).
		First(&user).
		Error

	return
}

func FindByID(id uint64) (user User, err error) {

	err = database.DB.Where("id", id).
		First(&user).
		Error

	return
}
