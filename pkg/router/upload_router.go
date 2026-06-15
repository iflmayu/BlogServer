package router

import (
	"BlogServer/internal/upload/handler"
	"BlogServer/internal/upload/repo"
	"BlogServer/internal/upload/service"
	"BlogServer/pkg/config"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerUploadRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config, jwtService *jwt.Service) {
	var uploader storage.Uploader

	switch cfg.Storage.Type {
	case "qiniu":
		q := cfg.Storage.Qiniu
		uploader = storage.NewQiniuStorage(q.AccessKey, q.SecretKey, q.Bucket, q.Domain)
	default:
		l := cfg.Storage.Local
		uploader = storage.NewLocalStorage(l.RootPath, l.BaseURL)
	}

	uploadRepo := repo.NewUploadRepo(db)
	uploadService := service.NewUploadService(uploadRepo, uploader, cfg.Upload)
	uploadHandler := handler.NewUploadHandler(uploadService, jwtService)

	// 公开路由
	uploadHandler.RegisterRoutes(r)
}
