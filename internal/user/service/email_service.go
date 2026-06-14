package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"BlogServer/pkg/email"
	"BlogServer/pkg/redis"
)

type EmailService struct {
	emailService *email.Service
}

func NewEmailService(emailService *email.Service) *EmailService {
	return &EmailService{emailService: emailService}
}

func (s *EmailService) SendVerifyCode(ctx context.Context, toEmail string) (string, error) {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	subject := "邮箱验证码"
	body := fmt.Sprintf("您的验证码是：%s，5分钟内有效，请勿泄露给他人。", code)

	if err := s.emailService.Send(toEmail, subject, body); err != nil {
		return "", err
	}

	// 存入 Redis，5 分钟过期
	key := fmt.Sprintf("email:code:%s", toEmail)
	redis.Client.Set(ctx, key, code, 5*time.Minute)

	return code, nil
}

func (s *EmailService) VerifyCode(ctx context.Context, email, code string) bool {
	key := fmt.Sprintf("email:code:%s", email)
	storedCode, err := redis.Client.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	return storedCode == code
}
