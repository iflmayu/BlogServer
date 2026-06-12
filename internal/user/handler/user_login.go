package handler

import (
	"BlogServer/pkg/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	// 绑定并校验请求参数
	req := middleware.GetRequest[LoginRequest](c)

	token, err := h.userService.Login(c.Request.Context(), req.Username, req.Password)

	fmt.Println(token, err)

	//if err != nil {
	//	response.FailWithError(err, c)
	//}
}
