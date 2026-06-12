package router

import (
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"

	"BlogServer/internal/user/handler"

	"BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
)

func registerUserRoutes(r *gin.RouterGroup, jwtService *jwt.Service) {
	//userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(jwtService)
	userHandler := handler.NewUserHandler(userService)

	// 公开路由
	userHandler.RegisterRoutes(r)

	// 需要登录的路由
	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware(jwtService))
	{
		auth.GET("/detail", nil)
	}
}
