package pwd

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对明文密码进行加密
func HashPassword(password string) (string, error) {
	// GenerateFromPassword 会自动生成随机盐值并混入
	// DefaultCost 默认值为 10
	if password == "" {
		return "", errors.New("密码为空")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 校验明文密码与数据库中的密文是否匹配
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
