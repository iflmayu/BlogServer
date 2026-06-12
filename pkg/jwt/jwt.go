package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secret string
	issuer string
	expire time.Duration
}

func NewService(secret, issuer string, expire int) *Service {
	return &Service{
		secret: secret,
		issuer: issuer,
		expire: time.Duration(expire) * time.Hour,
	}
}

type Claims struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
}
type MyClaims struct {
	Claims
	jwt.RegisteredClaims // v5 推荐嵌套这个结构体处理过期时间等
}

// GenerateToken 生成 Token
func (s *Service) GenerateToken(claims Claims) (string, error) {
	claim := MyClaims{
		Claims: claims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expire)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.issuer, // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(s.secret))
}

func (s *Service) ParseToken(tokenString string) (*MyClaims, error) {
	if tokenString == "" {
		return nil, errors.New("请登录, token为空")
	}
	// 解析字符串并映射到 MyClaims 结构体
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 额外校验算法，防止算法混淆攻击
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 提供你的密钥，用于校验签名（必须和签发时一致）
		return []byte(s.secret), nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, errors.New("登录已过期")
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, errors.New("token 格式错误")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, errors.New("token 签名无效")
		default:
			return nil, errors.New("token 无效")
		}
	}

	// 校验并断言返回的数据
	// Valid 字段会检查过期时间 (exp) 和签发时间 (iat)
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
