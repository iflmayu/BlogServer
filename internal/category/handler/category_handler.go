package handler

import (
	"BlogServer/internal/category/service"
	userService "BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
	jwtService      *jwt.Service
	userService     *userService.UserService
}

func NewCategoryHandler(
	categoryService *service.CategoryService,
	jwtService *jwt.Service,
	userService *userService.UserService,
) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		jwtService:      jwtService,
		userService:     userService,
	}
}

func (h *CategoryHandler) RegisterRoutes(r *gin.RouterGroup) {
	admin := r.Group("/category")
	admin.Use(middleware.AdminMiddleware(h.jwtService, h.userService))
	{
		admin.POST("", middleware.BindJSON[CreateCategoryRequest](), h.CreateCategory)
	}
}
