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
		article.GET("", middleware.BindQuery[ListArticleRequest](), h.ListArticles)
		article.GET("/:id", middleware.BindUri[IDRequest](), h.GetArticle)
		article.POST("/:id/view", middleware.BindUri[IDRequest](), h.ViewArticle)
	}

	// 需要登录的路由
	auth := r.Group("/article")
	auth.Use(middleware.AuthMiddleware(h.jwtService))
	{
		auth.POST("/:id/like", middleware.BindUri[IDRequest](), h.LikeArticle)
	}

	// 需要管理员的路由
	admin := r.Group("/article")
	admin.Use(middleware.AdminMiddleware(h.jwtService, h.userService))
	{
		admin.POST("", middleware.BindJSON[CreateArticleRequest](), h.CreateArticle)
		admin.PUT("/:id", middleware.BindJSON[UpdateArticleRequest](), h.UpdateArticle)
		admin.DELETE("/:id", middleware.BindUri[IDRequest](), h.DeleteArticle)
	}
}
