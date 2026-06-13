package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"BlogServer/pkg/config"
)

var Client *redis.Client

func Init(cfg config.Redis) {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := Client.Ping(context.Background()).Err(); err != nil {
		zap.S().Fatalw("Redis 连接失败", "error", err)
	}

	zap.S().Infow("Redis 连接成功", "addr", cfg.Addr())
}
