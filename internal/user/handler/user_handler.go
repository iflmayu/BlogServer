package handler

import (
	uploadSvc "BlogServer/internal/upload/service"
	"BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户相关的 HTTP 接口处理
type UserHandler struct {
	userService   *service.UserService
	emailService  *service.EmailService
	jwtService    *jwt.Service
	uploadService *uploadSvc.UploadService
}

// NewUserHandler 创建 UserHandler
func NewUserHandler(
	userService *service.UserService,
	jwtService *jwt.Service,
	emailService *service.EmailService,
	uploadService *uploadSvc.UploadService,
) *UserHandler {
	return &UserHandler{
		userService:   userService,
		jwtService:    jwtService,
		emailService:  emailService,
		uploadService: uploadService,
	}
}

// RegisterRoutes 注册用户模块路由
func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	// 公开路由
	userGroup := r.Group("/user")
	{
		userGroup.GET("/captcha", h.GenerateCaptcha)
		userGroup.POST("/email/code/register", middleware.BindJSON[SendRegisterCodeRequest](), h.SendRegisterCode)
		userGroup.POST("/register/verify", middleware.BindJSON[VerifyRegisterRequest](), h.VerifyRegisterInfo)
		userGroup.POST("/register/complete", middleware.BindJSON[CompleteRegisterRequest](), h.CompleteRegister)
		userGroup.POST("/login/pwd", middleware.BindJSON[LoginRequest](), h.PwdLogin)
		userGroup.POST("/email/code/login", middleware.BindJSON[SendLoginCodeRequest](), h.SendLoginCode)
		userGroup.POST("/login/email", middleware.BindJSON[EmailLoginRequest](), h.LoginByEmail)
		userGroup.POST("/email/code/reset", middleware.BindJSON[SendResetPasswordCodeRequest](), h.SendResetPasswordCode)
		userGroup.POST("/password/reset", middleware.BindJSON[ResetPasswordRequest](), h.ResetPassword)
	}

	// 需要登录的路由
	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware(h.jwtService))
	{
		auth.POST("/avatar", h.UpdateAvatar)
		auth.GET("/profile", h.GetProfile)
		auth.POST("/email/code/bind", middleware.BindJSON[SendBindEmailCodeRequest](), h.SendBindEmailCode)
		auth.POST("/email/bind", middleware.BindJSON[BindEmailRequest](), h.BindEmail)
	}
}
