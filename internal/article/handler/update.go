package handler

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/article/service"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type IDRequest struct {
	ID uint `uri:"id"`
}
type UpdateArticleRequest struct {
	Title      string               `json:"title" binding:"required,max=256"`
	Abstract   string               `json:"abstract" binding:"max=512"`
	Content    string               `json:"content" binding:"required"`
	Cover      string               `json:"cover" binding:"max=512"`
	CategoryID uint                 `json:"category_id"`
	Tags       domain.StringArray   `json:"tags"`
	Status     domain.ArticleStatus `json:"status" binding:"required"`
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	var idReq IDRequest
	if err := c.ShouldBindUri(&idReq); err != nil {
		response.FailWithMsg("无效的文章ID", c)
		return
	}

	req := middleware.GetRequest[UpdateArticleRequest](c)

	if err := h.articleService.Update(c.Request.Context(), service.UpdateArticleInput{
		ID:         idReq.ID,
		Title:      req.Title,
		Abstract:   req.Abstract,
		Content:    req.Content,
		Cover:      req.Cover,
		CategoryID: req.CategoryID,
		Tags:       req.Tags,
		Status:     req.Status,
	}); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("文章更新成功", c)
}
