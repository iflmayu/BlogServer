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

func (r *CategoryRepo) GetByID(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepo) GetByNameOrSlug(ctx context.Context, name, slug string, excludeID uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).
		Where("(name = ? OR slug = ?) AND id != ?", name, slug, excludeID).
		First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepo) Update(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}
