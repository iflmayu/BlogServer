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

// CommentItem 包含用户信息和 @ 用户信息的评论项
type CommentItem struct {
	domain.ArticleComment
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// ListByArticleID 查询文章评论（平铺列表）
func (r *CommentRepo) ListByArticleID(ctx context.Context, articleID uint, page, pageSize int) ([]CommentItem, int64, error) {
	var total int64
	if err := r.db.WithContext(ctx).Model(&domain.ArticleComment{}).
		Where("article_id = ? AND status = ?", articleID, domain.CommentStatusNormal).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var comments []CommentItem
	offset := (page - 1) * pageSize
	err := r.db.WithContext(ctx).Raw(`
		SELECT c.*, u.username, u.nickname, u.avatar
		FROM article_comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.article_id = ? AND c.status = ?
		ORDER BY c.created_at DESC
		LIMIT ? OFFSET ?
	`, articleID, domain.CommentStatusNormal, pageSize, offset).Scan(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

// GetByID 根据 ID 查询评论
func (r *CommentRepo) GetByID(ctx context.Context, id uint) (*domain.ArticleComment, error) {
	var comment domain.ArticleComment
	if err := r.db.WithContext(ctx).First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// Delete 删除评论并同步文章评论数
func (r *CommentRepo) Delete(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var comment domain.ArticleComment
		if err := tx.First(&comment, commentID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&comment).Error; err != nil {
			return err
		}

		if err := tx.Model(&domain.Article{}).
			Where("id = ?", comment.ArticleID).
			Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
			return err
		}

		return nil
	})
}
