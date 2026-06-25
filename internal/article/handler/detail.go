package handler

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type ArticleDetailResponse struct {
	ID           uint               `json:"id"`
	Title        string             `json:"title"`
	Abstract     string             `json:"abstract"`
	Content      string             `json:"content"`
	Cover        string             `json:"cover"`
	CategoryID   uint               `json:"category_id"`
	Tags         domain.StringArray `json:"tags"`
	Status       string             `json:"status"`
	ViewCount    int64              `json:"view_count"`
	LikeCount    int64              `json:"like_count"`
	CommentCount int64              `json:"comment_count"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	req := middleware.GetRequest[request.IDRequest](c)

	article, err := h.articleService.GetArticleDetail(c.Request.Context(), req.ID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(ArticleDetailResponse{
		ID:           article.ID,
		Title:        article.Title,
		Abstract:     article.Abstract,
		Content:      article.Content,
		Cover:        article.Cover,
		CategoryID:   article.CategoryID,
		Tags:         article.Tags,
		Status:       article.Status.String(),
		ViewCount:    article.ViewCount,
		LikeCount:    article.LikeCount,
		CommentCount: article.CommentCount,
		CreatedAt:    article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, c)
}
