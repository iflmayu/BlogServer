package router

import (
	"BlogServer/internal/article/handler"

	"github.com/gin-gonic/gin"
)

func registerArticleRoutes(r *gin.RouterGroup, h *handler.ArticleHandler) {
	h.RegisterRoutes(r)
}
