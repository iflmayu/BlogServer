package handler

import (
	"BlogServer/internal/article/service"
	userService "BlogServer/internal/user/service"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService *service.ArticleService
	jwtService     *jwt.Service
	userService    *userService.UserService
}

func NewArticleHandler(
	articleService *service.ArticleService,
	jwtService *jwt.Service,
	userService *userService.UserService,
) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
		jwtService:     jwtService,
		userService:    userService,
	}
}

func (h *ArticleHandler) RegisterRoutes(r *gin.RouterGroup) {
	article := r.Group("/article")
	{
		article.GET("", h.ListArticles)
	}

	// 需要管理员的路由
	admin := r.Group("/article")
	admin.Use(middleware.AdminMiddleware(h.jwtService, h.userService))
	{
		admin.POST("", middleware.BindJSON[CreateArticleRequest](), h.CreateArticle)
	}
}
