package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UpdateAvatar(c *gin.Context) {
	// 上传头像
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		response.FailWithMsg("请选择要上传的图片", c)
		return
	}

	url, err := h.uploadService.UploadImage(c.Request.Context(), fileHeader)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// 安全取当前登录用户
	claims, exists := c.Get("claims")
	if !exists {
		response.FailWithMsg("请先登录", c)
		return
	}
	myClaims, ok := claims.(*jwt.MyClaims)
	if !ok {
		response.FailWithMsg("登录信息无效", c)
		return
	}

	// 更新 user.avatar
	err = h.userService.UpdateAvatar(c.Request.Context(), myClaims.UserID, url)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	}

	response.OkWithMsg("头像更换成功", c)
}
