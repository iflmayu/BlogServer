package handler

import (
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type ViewArticleResponse struct {
	ViewCount int64 `json:"view_count"`
}

func (h *ArticleHandler) ViewArticle(c *gin.Context) {
	req := middleware.GetRequest[request.IDRequest](c)

	viewCount, err := h.articleService.ViewArticle(c.Request.Context(), req.ID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(ViewArticleResponse{
		ViewCount: viewCount,
	}, c)
}
