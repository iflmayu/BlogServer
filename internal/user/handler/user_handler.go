package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户相关的 HTTP 接口处理
type UserHandler struct {
	userService  *service.UserService
	emailService *service.EmailService
	jwtService   *jwt.Service
}

// NewUserHandler 创建 UserHandler
func NewUserHandler(userService *service.UserService, jwtService *jwt.Service, emailService *service.EmailService) *UserHandler {
	return &UserHandler{
		userService:  userService,
		jwtService:   jwtService,
		emailService: emailService,
	}
}

// RegisterRoutes 注册用户模块路由
func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	// 公开路由
	userGroup := r.Group("/user")
	{
		userGroup.GET("/captcha", h.GenerateCaptcha)
		userGroup.POST("/register", middleware.BindJSON[RegisterRequest](), h.Register)
		userGroup.POST("/login", middleware.BindJSON[LoginRequest](), h.Login)
	}

	// 需要登录的路由
	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware(h.jwtService))
	{
		auth.GET("/detail", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			response.OkWithData(claims, c)
		})
		auth.POST("/email/code/bind", middleware.BindJSON[SendEmailCodeRequest](), h.SendEmailCode("bind"))
		auth.POST("/email/code/resetpassword", middleware.BindJSON[SendEmailCodeRequest](), h.SendEmailCode("reset_password"))
	}
}
