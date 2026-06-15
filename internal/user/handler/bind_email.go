package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

func (h *UserHandler) BindEmail(c *gin.Context) {
	req := middleware.GetRequest[BindEmailRequest](c)

	claims, _ := c.Get("claims")
	myClaims := claims.(*jwt.MyClaims)

	if err := h.userService.BindEmail(c.Request.Context(), myClaims.UserID, req.Email, req.Code); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("邮箱绑定成功", c)
}
