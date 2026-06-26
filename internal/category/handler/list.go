package handler

import (
	"BlogServer/internal/common/response"

	"github.com/gin-gonic/gin"
)

type ListCategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.categoryService.List(c.Request.Context())
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	list := make([]ListCategoryResponse, len(categories))
	for i, category := range categories {
		list[i] = ListCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Slug:        category.Slug,
			Description: category.Description,
			SortOrder:   category.SortOrder,
			CreatedAt:   category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   category.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	response.OkWithData(list, c)
}
