package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type ProfileResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		response.FailWithMsg("该接口需要登录", c)
		return
	}
	myClaims := claims.(*jwt.MyClaims)

	user, err := h.userService.GetByID(c.Request.Context(), myClaims.UserID)
	if err != nil {
		response.FailWithMsg("获取用户信息失败", c)
		return
	}

	response.OkWithData(ProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Role:      user.Role.String(),
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, c)
}
