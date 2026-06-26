package handler

import (
	"BlogServer/internal/category/service"
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required,max=32"`
	Slug        string `json:"slug" binding:"required,max=32"`
	Description string `json:"description" binding:"max=256"`
	SortOrder   int    `json:"sort_order"`
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var idReq request.IDRequest
	if err := c.ShouldBindUri(&idReq); err != nil {
		response.FailWithMsg("无效的类别ID", c)
		return
	}
	req := middleware.GetRequest[UpdateCategoryRequest](c)

	if err := h.categoryService.Update(c.Request.Context(), service.UpdateCategoryInput{
		ID:          idReq.ID,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("分类更新成功", c)
}
