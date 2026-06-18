package handler

import (
	"BlogServer/internal/upload/service"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadService *service.UploadService
	jwtService    *jwt.Service
}

// NewUploadHandler 创建 UploadHandler
func NewUploadHandler(uploadService *service.UploadService, jwtService *jwt.Service) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
		jwtService:    jwtService,
	}
}

// RegisterRoutes 注册 upload 模块路由
func (h *UploadHandler) RegisterRoutes(r *gin.RouterGroup) {
	// 需要登录的路由
	auth := r.Group("/upload")
	auth.Use(middleware.AuthMiddleware(h.jwtService))
	{
		auth.POST("/image", h.UploadImage)
	}
}
