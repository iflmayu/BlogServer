package router

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/email"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"

	"BlogServer/internal/user/handler"

	"BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
)

func registerUserRoutes(r *gin.RouterGroup, jwtService *jwt.Service, emailSvc *email.Service) {
	//userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(jwtService)
	emailService := service.NewEmailService(emailSvc)
	userHandler := handler.NewUserHandler(userService, emailService)

	// 公开路由
	userHandler.RegisterRoutes(r)

	// 需要登录的路由
	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware(jwtService))
	{
		auth.GET("/detail", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			response.OkWithData(claims, c)
		})
		auth.POST("/email", middleware.BindJSON[handler.SendEmailCodeRequest](), userHandler.SendEmailCode)
	}
}
