package service

import (
	"errors"
	"jigsaw-api/internal/model"
	"jigsaw-api/internal/repository"
	"jigsaw-api/pkg/utils"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// FindUsers 分页获取用户列表
func (s *UserService) FindUsers(page, pageSize int) ([]model.User, int64, error) {
	return s.userRepo.FindUsers(page, pageSize)
}

// GetUser 根据ID获取单个用户
func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.userRepo.GetUser(id)
}

// CreateUser 创建新用户
// 检查用户名和邮箱是否存在，对密码进行加密
func (s *UserService) CreateUser(user *model.User) error {
	if s.userRepo.CheckUsernameExist(user.Username) {
		return errors.New("用户名已存在")
	}
	if s.userRepo.CheckEmailExist(user.Email) {
		return errors.New("邮箱已存在")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.Create(user)
}

// UpdateUser 更新用户信息
// 支持更新头像和邮箱，如果更新邮箱则检查唯一性
func (s *UserService) UpdateUser(userID uint, updates map[string]interface{}) error {
	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	if avatar, ok := updates["avatar"].(string); ok && avatar != "" {
		user.Avatar = avatar
	}

	if email, ok := updates["email"].(string); ok && email != "" {
		// If email is being changed, check if it's already taken by another user
		if email != user.Email {
			if s.userRepo.CheckEmailExist(email) {
				return errors.New("邮箱已存在")
			}
			user.Email = email
		}
	}

	return s.userRepo.Update(user)
}

func (s *UserService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		return errors.New("用户不存在")
	}
	if !utils.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("旧密码不正确")
	}
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userRepo.Update(user)
}
