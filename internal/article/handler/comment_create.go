package handler

import (
	"BlogServer/internal/article/service"
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/jwt"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	AtID    uint   `json:"at_id"`
	Content string `json:"content" binding:"required"`
}

type CreateCommentResponse struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	UserID    uint   `json:"user_id"`
	AtID      uint   `json:"at_id"`
	Content   string `json:"content,max=1000"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (h *ArticleHandler) CreateComment(c *gin.Context) {
	var idReq request.IDRequest
	if err := c.ShouldBindUri(&idReq); err != nil {
		response.FailWithMsg("无效的文章ID", c)
		return
	}

	req := middleware.GetRequest[CreateCommentRequest](c)

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

	comment, err := h.commentService.Create(c.Request.Context(), service.CreateCommentInput{
		ArticleID: idReq.ID,
		UserID:    myClaims.UserID,
		AtID:      req.AtID,
		Content:   req.Content,
	})
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(CreateCommentResponse{
		ID:        comment.ID,
		ArticleID: comment.ArticleID,
		UserID:    comment.UserID,
		AtID:      comment.AtID,
		Content:   comment.Content,
		Status:    comment.Status.String(),
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: comment.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, c)
}
