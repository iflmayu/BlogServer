package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type SendResetPasswordCodeRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (h *UserHandler) SendResetPasswordCode(c *gin.Context) {
	req := middleware.GetRequest[SendResetPasswordCodeRequest](c)

	// 校验图形验证码
	if !captcha.Verify(req.CaptchaID, req.CaptchaCode) {
		response.FailWithMsg("图形验证码错误", c)
		return
	}

	// 发送邮箱验证码
	if err := h.userService.SendResetPasswordCode(c.Request.Context(), req.Email); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("重置密码验证码已发送", c)
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	EmailCode   string `json:"email_code" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	req := middleware.GetRequest[ResetPasswordRequest](c)

	if err := h.userService.ResetPassword(
		c.Request.Context(),
		req.Email,
		req.EmailCode,
		req.NewPassword,
	); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("密码重置成功", c)
}
