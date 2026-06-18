package handler

import (
	"BlogServer/internal/common/response"

	"github.com/gin-gonic/gin"
)

func (h *UploadHandler) UploadImage(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.FailWithMsg("请选择要上传的文件", c)
		return
	}
	
	url, err := h.uploadService.UploadImage(c.Request.Context(), fileHeader)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{"url": url}, c)
}
