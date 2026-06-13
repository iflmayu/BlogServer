package router

import (
	"BlogServer/internal/upload/handler"
	"BlogServer/internal/upload/service"
	"BlogServer/pkg/config"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func registerUploadRoutes(r *gin.RouterGroup, cfg config.System, jwtService *jwt.Service) {
	baseURL := fmt.Sprintf("http://%s:%d/uploads", cfg.Ip, cfg.Port)
	uploadService := service.NewUploadService("./uploads", baseURL)
	uploadHandler := handler.NewUploadHandler(uploadService)

	// 公开路由
	uploadHandler.RegisterRoutes(r)

	// 需要登录的路由
	auth := r.Group("/upload")
	auth.Use(middleware.AuthMiddleware(jwtService))
	{
		auth.POST("/images", uploadHandler.UploadImage)
	}

}
