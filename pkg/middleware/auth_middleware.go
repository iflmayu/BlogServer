package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"
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
