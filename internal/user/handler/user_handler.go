package handler

import (
	"BlogServer/internal/user/service"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户相关的 HTTP 接口处理
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler 创建 UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// RegisterRoutes 注册用户模块路由
func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", middleware.BindJSON[LoginRequest](), h.Login)
	}
}
