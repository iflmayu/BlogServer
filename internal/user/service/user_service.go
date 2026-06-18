package service

import (
	"BlogServer/internal/user/domain"
	"BlogServer/internal/user/repo"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/redis"
	"BlogServer/pkg/utils/pwd"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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

const registerTokenTTL = 5 * time.Minute

func (s *UserService) SendRegisterCode(ctx context.Context, email string) error {
	_, err := s.emailService.SendVerifyCode(ctx, "register", email)
	return err
}

func (s *UserService) CheckEmailAvailable(ctx context.Context, email string) error {
	exists, err := s.userRepo.ExistsEmail(ctx, email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("该邮箱已被注册")
	}
	return nil
}

func (s *UserService) VerifyRegisterInfo(ctx context.Context, email, username, code string) (string, error) {
	// 校验邮箱验证码
	if !s.emailService.VerifyCode(ctx, "register", email, code) {
		return "", errors.New("邮箱验证码错误或已过期")
	}

	// 检查用户名
	exists, err := s.userRepo.ExistsUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("用户名已存在")
	}

	// 再次检查邮箱
	if err := s.CheckEmailAvailable(ctx, email); err != nil {
		return "", err
	}

	// 生成临时 token
	registerToken := uuid.New().String()
	key := fmt.Sprintf("register:token:%s", registerToken)
	value := fmt.Sprintf("%s|%s", email, username)
	redis.Client.Set(ctx, key, value, registerTokenTTL)

	return registerToken, nil
}

func (s *UserService) CompleteRegister(ctx context.Context, registerToken, password, avatar string) error {
	key := fmt.Sprintf("register:token:%s", registerToken)
	value, err := redis.Client.Get(ctx, key).Result()
	if err != nil {
		return errors.New("注册令牌无效或已过期")
	}

	parts := strings.SplitN(value, "|", 2)
	if len(parts) != 2 {
		return errors.New("注册令牌数据异常")
	}
	email := parts[0]
	username := parts[1]

	hashedPassword, err := pwd.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Avatar:   avatar,
		Role:     domain.RoleUser,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	redis.Client.Del(ctx, key)
	return nil
}

func (s *UserService) PwdLogin(ctx context.Context, username, password string) (string, error) {
	// 查用户
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 校验密码
	if !pwd.CheckPassword(user.Password, password) {
		return "", errors.New("用户名或密码错误")
	}

	// 生成 JWT
	token, err := s.jwtService.GenerateToken(jwt.Claims{
		UserID:   user.ID,
		Username: user.Username,
	})
	if err != nil {
		return "", errors.New("token 生成失败")
	}

	return token, nil
}

func (s *UserService) SendLoginCode(ctx context.Context, email string) error {
	// 检查邮箱是否已注册
	exists, err := s.userRepo.ExistsEmail(ctx, email)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("该邮箱未注册")
	}

	_, err = s.emailService.SendVerifyCode(ctx, "login", email)
	return err
}

func (s *UserService) LoginByEmail(ctx context.Context, email, code string) (string, error) {
	// 校验邮箱验证码
	if !s.emailService.VerifyCode(ctx, "login", email, code) {
		return "", errors.New("邮箱验证码错误或已过期")
	}

	// 查用户
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("该邮箱未注册")
	}

	// 生成 JWT
	token, err := s.jwtService.GenerateToken(jwt.Claims{
		UserID:   user.ID,
		Username: user.Username,
	})
	if err != nil {
		return "", errors.New("token 生成失败")
	}

	return token, nil
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
