package router

import (
	"BlogServer/pkg/config"
	"BlogServer/pkg/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config, jwtService *jwt.Service) *gin.Engine {
	gin.SetMode(cfg.System.GinMode)
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	api := r.Group("/api")

	// 注册 user 模块路由
	registerUserRoutes(api, jwtService)
	//
	registerUploadRoutes(api, db, cfg, jwtService)

	return r
}
