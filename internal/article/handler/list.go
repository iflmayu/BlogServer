package handler

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/article/service"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type ListArticleRequest struct {
	Page       int    `form:"page"`
	PageSize   int    `form:"page_size"`
	Keyword    string `form:"keyword"`
	CategoryID uint   `form:"category_id"`
}

type ListArticleResponse struct {
	ID           uint               `json:"id"`
	Title        string             `json:"title"`
	Abstract     string             `json:"abstract"`
	Cover        string             `json:"cover"`
	CategoryID   uint               `json:"category_id"`
	Tags         domain.StringArray `json:"tags"`
	Status       string             `json:"status"`
	ViewCount    int64              `json:"view_count"`
	LikeCount    int64              `json:"like_count"`
	CommentCount int64              `json:"comment_count"`
	CreatedAt    string             `json:"created_at"`
}

func (h *ArticleHandler) ListArticles(c *gin.Context) {
	req := middleware.GetRequest[ListArticleRequest](c)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 9
	}

	articles, total, err := h.articleService.List(c.Request.Context(), service.ListArticleInput{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Keyword:    req.Keyword,
		CategoryID: req.CategoryID,
		Status:     domain.ArticleStatusPublished,
	})
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	list := make([]ListArticleResponse, len(articles))
	for i, article := range articles {
		list[i] = ListArticleResponse{
			ID:           article.ID,
			Title:        article.Title,
			Abstract:     article.Abstract,
			Cover:        article.Cover,
			CategoryID:   article.CategoryID,
			Tags:         article.Tags,
			Status:       article.Status.String(),
			ViewCount:    article.ViewCount,
			LikeCount:    article.LikeCount,
			CommentCount: article.CommentCount,
			CreatedAt:    article.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	response.OkWithList(list, int(total), c)
}
