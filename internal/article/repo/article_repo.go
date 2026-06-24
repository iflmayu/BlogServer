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

type ListArticleQuery struct {
	Page       int
	PageSize   int
	Keyword    string
	CategoryID uint
	Status     domain.ArticleStatus
}

func (r *ArticleRepo) List(ctx context.Context, query *ListArticleQuery) ([]domain.Article, int64, error) {
	db := r.db.WithContext(ctx).Model(&domain.Article{})

	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}
	if query.CategoryID != 0 {
		db = db.Where("category_id = ?", query.CategoryID)
	}
	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"
		db = db.Where("title ILIKE ? OR abstract ILIKE ?", keyword, keyword)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var articles []domain.Article
	offset := (query.Page - 1) * query.PageSize
	err := db.Order("created_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&articles).Error

	return articles, total, err
}
