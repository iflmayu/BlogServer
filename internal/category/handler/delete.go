package handler

import (
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idReq := middleware.GetRequest[request.IDRequest](c)

	if err := h.categoryService.Delete(c.Request.Context(), idReq.ID); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("分类删除成功", c)
}
