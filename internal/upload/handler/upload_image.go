package handler

import (
	"BlogServer/internal/common/response"

	"github.com/gin-gonic/gin"
)

func (h *UploadHandler) UploadImage(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	//url, err := h.uploadService.UploadImage(fileHeader)
	//if err != nil {
	//	response.FailWithMsg(err.Error(), c)
	//	return
	//}

	response.OkWithData(gin.H{"fileHeader": fileHeader}, c)
}
