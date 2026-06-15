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

const emailCodeTTL = 5 * time.Minute

var subjectMap = map[string]string{
	"bind":           "邮箱绑定验证码",
	"reset_password": "重置密码验证码",
}

func getSubject(action string) string {
	if subject, ok := subjectMap[action]; ok {
		return subject
	}
	return "邮箱验证码"
}

func (s *EmailService) SendVerifyCode(ctx context.Context, action, toEmail string) (string, error) {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	subject := getSubject(action)
	body := fmt.Sprintf("您的验证码是：%s，%d分钟内有效，请勿泄露给他人。", code, int(emailCodeTTL.Minutes()))

	if err := s.emailService.Send(toEmail, subject, body); err != nil {
		return "", err
	}

	// 存入 Redis，5 分钟过期
	key := fmt.Sprintf("email:code:%s:%s", toEmail, action)
	redis.Client.Set(ctx, key, code, 5*time.Minute)

	return code, nil
}

func (s *EmailService) VerifyCode(ctx context.Context, action, email, code string) bool {
	key := fmt.Sprintf("email:code:%s:%s", email, action)
	storedCode, err := redis.Client.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if storedCode != code {
		return false
	}
	redis.Client.Del(ctx, key)
	return true
}
