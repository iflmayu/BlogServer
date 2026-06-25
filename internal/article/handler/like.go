package handler

import (
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type LikeArticleRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type LikeArticleResponse struct {
	LikeCount int64 `json:"like_count"`
	IsLiked   bool  `json:"is_liked"`
}

func (h *ArticleHandler) LikeArticle(c *gin.Context) {
	req := middleware.GetRequest[IDRequest](c)

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

	isLiked, likeCount, err := h.articleService.ToggleLike(
		c.Request.Context(),
		req.ID,
		myClaims.UserID,
	)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(LikeArticleResponse{
		LikeCount: likeCount,
		IsLiked:   isLiked,
	}, c)
}
