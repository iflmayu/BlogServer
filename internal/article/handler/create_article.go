package handler

import (
	"BlogServer/internal/article/service"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type CreateArticleRequest struct {
	Title      string   `json:"title" binding:"required,max=256"`
	Abstract   string   `json:"abstract" binding:"max=512"`
	Content    string   `json:"content" binding:"required"`
	Cover      string   `json:"cover" binding:"max=512"`
	CategoryID uint     `json:"category_id"`
	Tags       []string `json:"tags"`
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	req := middleware.GetRequest[CreateArticleRequest](c)

	input := service.CreateArticleInput{
		Title:      req.Title,
		Abstract:   req.Abstract,
		Content:    req.Content,
		Cover:      req.Cover,
		CategoryID: req.CategoryID,
		Tags:       req.Tags,
	}

	if err := h.articleService.Create(c.Request.Context(), input); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("文章发布成功", c)
}
