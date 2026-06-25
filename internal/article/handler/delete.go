package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type DeleteArticleRequest struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	req := middleware.GetRequest[IDRequest](c)

	if err := h.articleService.Delete(c.Request.Context(), req.ID); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("文章删除成功", c)
}
