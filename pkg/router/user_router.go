package router

import (
	"BlogServer/internal/user/handler"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(r *gin.RouterGroup, h *handler.UserHandler) {
	h.RegisterRoutes(r)
}
