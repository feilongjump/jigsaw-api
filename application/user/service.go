package user

import (
	"errors"

	"github.com/feilongjump/jigsaw-api/application/user/dto"
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepo repo.UserRepository
}

func NewService(userRepo repo.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

// Register 用户注册
func (s *Service) Register(req *dto.RegisterRequest) error {
	// 1. 检查用户是否存在
	existUser, err := s.userRepo.FindByUsername(req.Username)
	// 查询当前用户是否存在，所以需要增加判断是否是用户不存在的错误
	if err != nil && !errors.Is(err, err_code.UserNotFound) {
		return err
	}
	if existUser != nil {
		return err_code.UserAlreadyExists
	}

	// 2. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. 创建用户
	user := &entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}

// Login 用户登录
func (s *Service) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. 查找用户
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, err_code.UserNotFound
	}

	// 2. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err_code.UserPasswordError
	}

	// 3. 生成 Token
	token, err := jwt.GenerateToken(uint64(user.ID))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}

// ChangePassword 修改密码
func (s *Service) ChangePassword(userID uint64, req *dto.ChangePasswordRequest) error {
	// 1. 查找用户
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return err_code.UserNotFound
	}

	// 2. 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return err_code.UserPasswordError
	}

	// 3. 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 4. 更新密码
	if err := s.userRepo.UpdatePassword(userID, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}
