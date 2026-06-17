package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PwdLogin 用户登录
func (h *UserHandler) PwdLogin(c *gin.Context) {
	req := middleware.GetRequest[LoginRequest](c)

	token, err := h.userService.PwdLogin(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{"token": token}, c)
}

type SendLoginCodeRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (h *UserHandler) SendLoginCode(c *gin.Context) {
	req := middleware.GetRequest[SendLoginCodeRequest](c)

	// 校验图形验证码
	if !captcha.Verify(req.CaptchaID, req.CaptchaCode) {
		response.FailWithMsg("图形验证码错误", c)
		return
	}

	if err := h.userService.SendLoginCode(c.Request.Context(), req.Email); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("登录验证码已发送", c)
}

type EmailLoginRequest struct {
	Email     string `json:"email" binding:"required,email"`
	EmailCode string `json:"email_code" binding:"required"`
}

func (h *UserHandler) LoginByEmail(c *gin.Context) {
	req := middleware.GetRequest[EmailLoginRequest](c)

	token, err := h.userService.LoginByEmail(c.Request.Context(), req.Email, req.EmailCode)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{"token": token}, c)
}
