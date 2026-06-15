package service

import (
	"BlogServer/internal/user/domain"
	"BlogServer/internal/user/repo"
	"BlogServer/pkg/jwt"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo     *repo.UserRepo
	jwtService   *jwt.Service
	emailService *EmailService
}

func NewUserService(userRepo *repo.UserRepo, jwtService *jwt.Service, emailService *EmailService) *UserService {
	return &UserService{
		userRepo:     userRepo,
		jwtService:   jwtService,
		emailService: emailService,
	}
}

func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	//TODO 数据库查询

	token, err := s.jwtService.GenerateToken(jwt.Claims{
		UserID:   0,
		Username: "xxx",
	})
	return token, err
}

func (s *UserService) BindEmail(ctx context.Context, userID uint, email, code string) error {
	// 校验邮箱验证码
	if !s.emailService.VerifyCode(ctx, "bind", email, code) {
		return errors.New("邮箱验证码错误或已过期")
	}

	return nil
	//// 检查邮箱是否已被其他账号绑定
	//exists, err := s.userRepo.ExistsEmail(ctx, email)
	//if err != nil {
	//	return err
	//}
	//if exists {
	//	return errors.New("该邮箱已被绑定")
	//}
	//
	//// 更新用户邮箱
	//return s.userRepo.UpdateEmail(ctx, userID, email)
}

func (s *UserService) Register(ctx context.Context, username, password string) error {
	// 检查用户名是否存在
	exists, err := s.userRepo.ExistsUsername(ctx, username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 创建用户
	user := &domain.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     domain.RoleUser,
	}

	return s.userRepo.Create(ctx, user)
}
