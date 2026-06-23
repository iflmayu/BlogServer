package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type SendRegisterCodeRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (h *UserHandler) SendRegisterCode(c *gin.Context) {
	req := middleware.GetRequest[SendRegisterCodeRequest](c)

	// 校验图形验证码
	if !captcha.Verify(req.CaptchaID, req.CaptchaCode) {
		response.FailWithMsg("图形验证码错误", c)
		return
	}

	// 检查邮箱是否已注册
	if err := h.userService.CheckEmailAvailable(c.Request.Context(), req.Email); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// 发送邮箱验证码
	if err := h.userService.SendRegisterCode(c.Request.Context(), req.Email); err != nil {
		response.FailWithMsg("邮件发送失败", c)
		return
	}

	response.OkWithMsg("注册验证码已发送", c)
}

type VerifyRegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required,min=3,max=20"`
	EmailCode string `json:"email_code" binding:"required"`
}

func (h *UserHandler) VerifyRegisterInfo(c *gin.Context) {
	req := middleware.GetRequest[VerifyRegisterRequest](c)

	token, err := h.userService.VerifyRegisterInfo(
		c.Request.Context(),
		req.Email,
		req.Username,
		req.EmailCode,
	)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{"register_token": token}, c)
}

type CompleteRegisterRequest struct {
	RegisterToken string `json:"register_token" binding:"required"`
	Password      string `json:"password" binding:"required,min=8"`
}

func (h *UserHandler) CompleteRegister(c *gin.Context) {
	req := middleware.GetRequest[CompleteRegisterRequest](c)

	if err := h.userService.CompleteRegister(
		c.Request.Context(),
		req.RegisterToken,
		req.Password,
	); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("注册成功", c)
}
