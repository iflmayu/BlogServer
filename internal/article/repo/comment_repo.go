package repo

import (
	"BlogServer/internal/article/domain"
	"context"

	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

// Create 创建评论并同步文章评论数
func (r *CommentRepo) Create(ctx context.Context, comment *domain.ArticleComment) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}

		if err := tx.Model(&domain.Article{}).
			Where("id = ?", comment.ArticleID).
			Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			return err
		}

		return nil
	})
}
