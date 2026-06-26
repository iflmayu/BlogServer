package router

import (
	articleHandler "BlogServer/internal/article/handler"
	articleRepo "BlogServer/internal/article/repo"
	articleService "BlogServer/internal/article/service"
	categoryHandler "BlogServer/internal/category/handler"
	categoryRepo "BlogServer/internal/category/repo"
	categoryService "BlogServer/internal/category/service"
	uploadHandler "BlogServer/internal/upload/handler"
	uploadRepo "BlogServer/internal/upload/repo"
	uploadService "BlogServer/internal/upload/service"
	userHandler "BlogServer/internal/user/handler"
	userRepo "BlogServer/internal/user/repo"
	userService "BlogServer/internal/user/service"
	"BlogServer/pkg/config"
	"BlogServer/pkg/email"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config, jwtService *jwt.Service, emailService *email.Service) *gin.Engine {
	gin.SetMode(cfg.System.GinMode)
	r := gin.Default()
	r.Static("/api/storage/image", cfg.Upload.UploadDir)

	api := r.Group("/api")

	// 创建所有 Service
	uRepo := userRepo.NewUserRepo(db)
	emailSvc := userService.NewEmailService(emailService)
	uSvc := userService.NewUserService(uRepo, jwtService, emailSvc)

	upRepo := uploadRepo.NewUploadRepo(db)
	upSvc := uploadService.NewUploadService(upRepo, newUploader(cfg), cfg.Upload)

	aRepo := articleRepo.NewArticleRepo(db)
	articleSvc := articleService.NewArticleService(aRepo)

	cRepo := articleRepo.NewCommentRepo(db)
	commentSvc := articleService.NewCommentService(cRepo, aRepo, uSvc)

	catRepo := categoryRepo.NewCategoryRepo(db)
	catSvc := categoryService.NewCategoryService(catRepo)

	// 创建所有 Handler
	uHandler := userHandler.NewUserHandler(uSvc, jwtService, emailSvc, upSvc)
	upHandler := uploadHandler.NewUploadHandler(upSvc, jwtService, uSvc)
	aHandler := articleHandler.NewArticleHandler(articleSvc, commentSvc, jwtService, uSvc)
	catHandler := categoryHandler.NewCategoryHandler(catSvc, jwtService, uSvc)

	// 注册路由
	registerUserRoutes(api, uHandler)
	registerUploadRoutes(api, upHandler)
	registerArticleRoutes(api, aHandler)
	registerCategoryRoutes(api, catHandler)

	return r
}

func newUploader(cfg *config.Config) storage.Uploader {
	switch cfg.Storage.Type {
	case "qiniu":
		q := cfg.Storage.Qiniu
		return storage.NewQiniuStorage(q.AccessKey, q.SecretKey, q.Bucket, q.Domain)
	default:
		return storage.NewLocalStorage(cfg.Storage.Local.BaseURL)
	}
}
