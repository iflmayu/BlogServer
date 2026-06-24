package repo

import (
	"BlogServer/internal/article/domain"
	"context"

	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{db: db}
}

func (r *ArticleRepo) Create(ctx context.Context, article *domain.Article) error {
	return r.db.WithContext(ctx).Create(article).Error
}
