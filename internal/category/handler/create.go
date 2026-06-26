package handler

import (
	"BlogServer/internal/category/service"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required,max=32"`
	Slug        string `json:"slug" binding:"required,max=32"`
	Description string `json:"description" binding:"max=256"`
	SortOrder   int    `json:"sort_order"`
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	req := middleware.GetRequest[CreateCategoryRequest](c)

	if err := h.categoryService.Create(c.Request.Context(), service.CreateCategoryInput{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("分类创建成功", c)
}
