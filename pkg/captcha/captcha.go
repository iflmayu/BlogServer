package captcha

import (
	"context"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"

	"BlogServer/pkg/config"
)

var instance *base64Captcha.Captcha

func Init(cfg config.Captcha, redisClient *redis.Client) {
	driver := base64Captcha.NewDriverDigit(
		cfg.Height,
		cfg.Width,
		cfg.Length,
		0.7,
		80,
	)
	store := newRedisStore(redisClient, time.Duration(cfg.ExpireSeconds)*time.Second)
	instance = base64Captcha.NewCaptcha(driver, store)
}

func Generate() (id, b64s string, err error) {
	id, b64s, _, err = instance.Generate()
	return
}

func Verify(id, answer string) bool {
	return instance.Verify(id, answer, true)
}

type redisStore struct {
	client *redis.Client
	ttl    time.Duration
	prefix string
}

func newRedisStore(client *redis.Client, ttl time.Duration) *redisStore {
	return &redisStore{
		client: client,
		ttl:    ttl,
		prefix: "captcha:",
	}
}

func (s *redisStore) Set(id string, value string) error {
	return s.client.Set(context.Background(), s.prefix+id, value, s.ttl).Err()
}

func (s *redisStore) Get(id string, clear bool) string {
	key := s.prefix + id
	val, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.client.Del(context.Background(), key)
	}
	return val
}

func (s *redisStore) Verify(id, answer string, clear bool) bool {
	val := s.Get(id, clear)
	return val != "" && val == answer
}
