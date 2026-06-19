package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/redis"
)

func AuthMiddleware(jwtService *jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.FailWithMsg("请先登录", c)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.FailWithMsg("token 格式错误", c)
			c.Abort()
			return
		}

		key := fmt.Sprintf("register:token:%s", parts[1])
		_, err := redis.Client.Get(c, key).Result()
		if err == nil {
			c.Set("register_token", parts[1])
			c.Next()
			return
		}

		if jwtService.IsBlacklisted(parts[1]) {
			fmt.Println(parts[1])
			response.FailWithMsg("token 黑名单", c)
			c.Abort()
			return
		}

		claims, err := jwtService.ParseToken(parts[1])
		if err != nil {
			response.FailWithMsg("登录已过期或 token 无效", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
