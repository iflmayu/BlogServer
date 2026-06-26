package router

import (
	"BlogServer/internal/category/handler"

	"github.com/gin-gonic/gin"
)

func registerCategoryRoutes(r *gin.RouterGroup, h *handler.CategoryHandler) {
	h.RegisterRoutes(r)
}
