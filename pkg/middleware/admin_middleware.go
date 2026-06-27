package middleware

import (
	"github.com/gin-gonic/gin"

	"BlogServer/internal/common/response"
	userService "BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
)

func AdminMiddleware(jwtService *jwt.Service, userService *userService.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := authenticate(c, jwtService)
		if !ok {
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
