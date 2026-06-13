package jwt

import (
	"context"
	"time"

	"BlogServer/pkg/redis"
)

const tokenBlacklistPrefix = "token:blacklist:"

// Blacklist 把 token 加入黑名单，TTL 为 token 剩余有效期
func (s *Service) Blacklist(tokenString string) error {
	claims, err := s.ParseToken(tokenString)
	if err != nil {
		return err
	}

	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		return nil // token 已过期，不用加黑名单
	}

	return redis.Client.Set(context.Background(), tokenBlacklistPrefix+tokenString, "1", ttl).Err()
}

// IsBlacklisted 检查 token 是否在黑名单
func (s *Service) IsBlacklisted(tokenString string) bool {
	n, err := redis.Client.Exists(context.Background(), tokenBlacklistPrefix+tokenString).Result()
	return err == nil && n > 0
}
