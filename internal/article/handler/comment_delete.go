package handler

import (
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (h *ArticleHandler) DeleteComment(c *gin.Context) {
	req := middleware.GetRequest[request.IDRequest](c)

	claims, exists := c.Get("claims")
	if !exists {
		response.FailWithMsg("请先登录", c)
		return
	}
	myClaims, ok := claims.(*jwt.MyClaims)
	if !ok {
		response.FailWithMsg("登录信息无效", c)
		return
	}

	isAdmin, err := h.userService.IsAdmin(c.Request.Context(), myClaims.UserID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	if err := h.commentService.Delete(c.Request.Context(), req.ID, myClaims.UserID, isAdmin); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("评论删除成功", c)
}
