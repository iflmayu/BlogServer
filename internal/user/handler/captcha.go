package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/captcha"

	"github.com/gin-gonic/gin"
)

type GenerateCaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (h *UserHandler) GenerateCaptcha(c *gin.Context) {
	id, b64s, err := captcha.Generate()
	if err != nil {
		response.FailWithMsg("验证码生成失败", c)
		return
	}

	response.OkWithData(GenerateCaptchaResponse{
		CaptchaID: id,
		Captcha:   b64s,
	}, c)
}
