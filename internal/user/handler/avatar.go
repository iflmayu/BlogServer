package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UpdateAvatar(c *gin.Context) {
	// 上传头像
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.FailWithMsg("请选择要上传的图片", c)
		return
	}

	url, err := h.uploadService.UploadImage(c.Request.Context(), fileHeader)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	}

	// 更新 user.avatar
	claims, _ := c.Get("claims")
	myClaims := claims.(*jwt.MyClaims)

	err = h.userService.UpdateAvatar(c.Request.Context(), myClaims.UserID, url)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	}

	response.OkWithMsg("头像更换成功", c)
}
