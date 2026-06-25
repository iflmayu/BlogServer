package handler

import (
	"BlogServer/internal/article/service"
	"BlogServer/internal/common/request"
	"BlogServer/internal/common/response"
	"BlogServer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type ListCommentRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type ListCommentResponse struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	UserID    uint   `json:"user_id"`
	AtID      uint   `json:"at_id"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
}

func (h *ArticleHandler) ListComments(c *gin.Context) {
	var idReq request.IDRequest
	if err := c.ShouldBindUri(&idReq); err != nil {
		response.FailWithMsg("无效的文章ID", c)
		return
	}
	req := middleware.GetRequest[ListCommentRequest](c)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 50
	}

	comments, total, err := h.commentService.List(c.Request.Context(), service.ListCommentInput{
		ArticleID: idReq.ID,
		Page:      req.Page,
		PageSize:  req.PageSize,
	})
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	list := make([]ListCommentResponse, len(comments))
	for i, comment := range comments {
		list[i] = ListCommentResponse{
			ID:        comment.ID,
			ArticleID: comment.ArticleID,
			UserID:    comment.UserID,
			AtID:      comment.AtID,
			Content:   comment.Content,
			Status:    comment.Status.String(),
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			Username:  comment.Username,
			Nickname:  comment.Nickname,
			Avatar:    comment.Avatar,
		}
	}

	response.OkWithList(list, int(total), c)
}
