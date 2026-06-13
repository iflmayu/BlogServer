package router

import (
	"BlogServer/pkg/config"
	"BlogServer/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.System, jwtService *jwt.Service) *gin.Engine {
	gin.SetMode(cfg.GinMode)
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	api := r.Group("/api")

	// 注册 user 模块路由
	registerUserRoutes(api, jwtService)
	//
	registerUploadRoutes(api, cfg, jwtService)

	return r
}
