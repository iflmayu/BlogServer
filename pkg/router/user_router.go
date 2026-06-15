package router

import (
	"BlogServer/internal/user/repo"
	"BlogServer/pkg/email"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"BlogServer/internal/user/handler"

	"BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
)

func registerUserRoutes(r *gin.RouterGroup, db *gorm.DB, jwtService *jwt.Service, emailSvc *email.Service) {
	userRepo := repo.NewUserRepo(db)
	emailService := service.NewEmailService(emailSvc)
	userService := service.NewUserService(userRepo, jwtService, emailService)
	userHandler := handler.NewUserHandler(userService, jwtService, emailService)

	userHandler.RegisterRoutes(r)
}
