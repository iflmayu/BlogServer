package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=20"`
	Password    string `json:"password" binding:"required,min=6"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
	req := middleware.GetRequest[RegisterRequest](c)

	if !captcha.Verify(req.CaptchaID, req.CaptchaCode) {
		response.FailWithMsg("图形验证码错误", c)
		return
	}

	err := h.userService.Register(
		c.Request.Context(),
		req.Username,
		req.Password,
	)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("注册成功", c)
}
