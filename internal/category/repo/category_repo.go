package repo

import (
	"BlogServer/internal/category/domain"
	"context"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Create(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *CategoryRepo) List(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.WithContext(ctx).
		Order("sort_order ASC, created_at DESC").
		Find(&categories).Error
	return categories, err
}
