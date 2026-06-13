package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/internal/upload/service"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadService *service.UploadService
}

// NewUploadHandler 创建 UploadHandler
func NewUploadHandler(uploadService *service.UploadService) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
	}
}

// RegisterRoutes 注册 upload 模块路由
func (h *UploadHandler) RegisterRoutes(r *gin.RouterGroup) {
	uploadGroup := r.Group("/upload")
	{
		uploadGroup.GET("", func(c *gin.Context) {
			response.OkWithData(gin.H{"upload": "upload"}, c)
		})
	}
}
