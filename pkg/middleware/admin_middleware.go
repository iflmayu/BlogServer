package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"BlogServer/internal/common/response"
	userService "BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
)

func AdminMiddleware(jwtService *jwt.Service, userService *userService.UserService) gin.HandlerFunc {
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

		isAdmin, err := userService.IsAdmin(c.Request.Context(), claims.UserID)
		if err != nil {
			response.FailWithMsg("用户信息查询失败", c)
			c.Abort()
			return
		}
		if !isAdmin {
			response.FailWithMsg("权限不足", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
