package repo

import (
	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
	FindByID(id uint64) (*entity.User, error)
	UpdatePassword(id uint64, password string) error
}
