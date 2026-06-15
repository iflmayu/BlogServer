package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type SendEmailCodeRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (h *UserHandler) SendEmailCode(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := middleware.GetRequest[SendEmailCodeRequest](c)

		if !captcha.Verify(req.CaptchaID, req.CaptchaCode) {
			response.FailWithMsg("图形验证码错误", c)
			return
		}

		_, err := h.emailService.SendVerifyCode(c.Request.Context(), action, req.Email)
		if err != nil {
			response.FailWithMsg("邮件发送失败", c)
			return
		}
		response.OkWithMsg("邮箱验证码已发送", c)
	}
}
