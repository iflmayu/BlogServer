package router

import (
	"BlogServer/internal/upload/handler"

	"github.com/gin-gonic/gin"
)

func registerUploadRoutes(r *gin.RouterGroup, h *handler.UploadHandler) {
	h.RegisterRoutes(r)
}
