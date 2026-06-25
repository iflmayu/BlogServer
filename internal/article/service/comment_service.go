package service

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/article/repo"
	"context"
	"errors"
)

type CommentService struct {
	commentRepo *repo.CommentRepo
	articleRepo *repo.ArticleRepo
}

func NewCommentService(commentRepo *repo.CommentRepo, articleRepo *repo.ArticleRepo) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
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
