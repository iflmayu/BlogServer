package service

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/article/repo"
	userService "BlogServer/internal/user/service"
	"context"
	"errors"
)

type CommentService struct {
	commentRepo *repo.CommentRepo
	articleRepo *repo.ArticleRepo
	userService *userService.UserService
}

func NewCommentService(commentRepo *repo.CommentRepo, articleRepo *repo.ArticleRepo, userService *userService.UserService) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
		userService: userService,
	}
}

type CreateCommentInput struct {
	ArticleID uint
	UserID    uint
	AtID      uint
	Content   string
}

func (s *CommentService) Create(ctx context.Context, input CreateCommentInput) (*domain.ArticleComment, error) {
	article, err := s.articleRepo.GetByID(ctx, input.ArticleID)
	if err != nil {
		return nil, wrapNotFound(err, "文章不存在")
	}
	if article.Status != domain.ArticleStatusPublished {
		return nil, errors.New("文章不存在或已下线")
	}

	// 校验 @ 的用户是否存在
	if input.AtID > 0 {
		if input.AtID == input.UserID {
			return nil, errors.New("不能@自己")
		}
		if _, err := s.userService.GetByID(ctx, input.AtID); err != nil {
			return nil, wrapNotFound(err, "@的用户不存在")
		}
	}

	comment := &domain.ArticleComment{
		ArticleID: input.ArticleID,
		UserID:    input.UserID,
		AtID:      input.AtID,
		Content:   input.Content,
		Status:    domain.CommentStatusNormal,
	}

	if err := s.commentRepo.Create(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

type ListCommentInput struct {
	ArticleID uint
	Page      int
	PageSize  int
}

func (s *CommentService) List(ctx context.Context, input ListCommentInput) ([]repo.CommentItem, int64, error) {
	article, err := s.articleRepo.GetByID(ctx, input.ArticleID)
	if err != nil {
		return nil, 0, wrapNotFound(err, "文章不存在")
	}
	if article.Status != domain.ArticleStatusPublished {
		return nil, 0, errors.New("文章不存在或已下线")
	}

	return s.commentRepo.ListByArticleID(ctx, input.ArticleID, input.Page, input.PageSize)
}

func (s *CommentService) Delete(ctx context.Context, commentID, userID uint, isAdmin bool) error {
	comment, err := s.commentRepo.GetByID(ctx, commentID)
	if err != nil {
		return wrapNotFound(err, "评论不存在")
	}

	if comment.UserID != userID && !isAdmin {
		return errors.New("无权删除该评论")
	}

	if err := s.commentRepo.Delete(ctx, commentID); err != nil {
		return wrapNotFound(err, "评论不存在")
	}

	return nil
}
