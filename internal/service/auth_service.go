package service

import (
	"errors"
	"strings"

	"jigsaw-api/internal/model"
	"jigsaw-api/internal/repository"
	"jigsaw-api/pkg/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}

// Register 用户注册逻辑
// 检查用户是否已存在，加密密码，创建用户
func (s *AuthService) Register(user *model.User) error {
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

// Login 用户登录逻辑
// 验证账号和密码，生成 JWT Token
func (s *AuthService) Login(account, password string) (string, error) {
	var user *model.User
	var err error
	if strings.Contains(account, "@") {
		user, err = s.userRepo.GetUserByEmail(account)
	} else {
		user, err = s.userRepo.GetUserByUsername(account)
	}
	if err != nil {
		return "", errors.New("账号或密码错误")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("账号或密码错误")
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
