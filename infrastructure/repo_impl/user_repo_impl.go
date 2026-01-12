package repo_impl

import (
	"errors"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository() repo.UserRepository {
	return &userRepositoryImpl{
		db: db.Get(),
	}
}

func (u *userRepositoryImpl) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

func (u *userRepositoryImpl) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.UserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindByID(id uint64) (*entity.User, error) {
	var user entity.User
	if err := u.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.UserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) UpdatePassword(id uint64, password string) error {
	return u.db.
		Model(&entity.User{
			ID: id,
		}).
		Update("password", password).
		Error
}

func (u *userRepositoryImpl) UpdateAvatar(id uint64, avatar string) error {
	return u.db.
		Model(&entity.User{
			ID: id,
		}).
		Update("avatar", avatar).
		Error
}
