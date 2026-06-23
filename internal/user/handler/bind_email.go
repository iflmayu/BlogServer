package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type SendBindEmailCodeRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (h *UserHandler) SendBindEmailCode(c *gin.Context) {
	req := middleware.GetRequest[SendBindEmailCodeRequest](c)

	// 校验图形验证码
	if !captcha.Verify(req.CaptchaID, req.CaptchaCode) {
		response.FailWithMsg("图形验证码错误", c)
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		response.FailWithMsg("请先登录", c)
		return
	}
	myClaims := claims.(*jwt.MyClaims)

	if err := h.userService.SendBindEmailCode(c.Request.Context(), myClaims.UserID, req.Email); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("邮箱绑定验证码已发送", c)
}

type BindEmailRequest struct {
	Email     string `json:"email" binding:"required,email"`
	EmailCode string `json:"email_code" binding:"required"`
}

func (h *UserHandler) BindEmail(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		response.FailWithMsg("请先登录", c)
		return
	}
	myClaims := claims.(*jwt.MyClaims)

	req := middleware.GetRequest[BindEmailRequest](c)

	if err := h.userService.BindEmail(c.Request.Context(), myClaims.UserID, req.Email, req.EmailCode); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("邮箱绑定成功", c)
}
