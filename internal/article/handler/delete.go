package handler

import (
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	req := middleware.GetRequest[request.IDRequest](c)

	if err := h.articleService.Delete(c.Request.Context(), req.ID); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("文章删除成功", c)
}
