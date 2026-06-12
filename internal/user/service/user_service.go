package service

import (
	"BlogServer/pkg/jwt"
	"context"
)

type UserService struct {
	jwtService *jwt.Service
}

func NewUserService(jwtService *jwt.Service) *UserService {
	return &UserService{
		jwtService: jwtService,
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
