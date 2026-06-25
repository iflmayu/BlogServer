package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type ViewArticleRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type ViewArticleResponse struct {
	ViewCount int64 `json:"view_count"`
}

func (h *ArticleHandler) ViewArticle(c *gin.Context) {
	req := middleware.GetRequest[IDRequest](c)

	viewCount, err := h.articleService.ViewArticle(c.Request.Context(), req.ID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(ViewArticleResponse{
		ViewCount: viewCount,
	}, c)
}
