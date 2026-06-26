package service

import (
	"BlogServer/internal/category/domain"
	"BlogServer/internal/category/repo"
	"context"
	"errors"
	"strings"
)

type CategoryService struct {
	categoryRepo *repo.CategoryRepo
}

func NewCategoryService(categoryRepo *repo.CategoryRepo) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

type CreateCategoryInput struct {
	Name        string
	Slug        string
	Description string
	SortOrder   int
}

func (s *CategoryService) Create(ctx context.Context, input CreateCategoryInput) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("分类名称不能为空")
	}
	if strings.TrimSpace(input.Slug) == "" {
		return errors.New("分类标识不能为空")
	}

	category := &domain.Category{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		SortOrder:   input.SortOrder,
	}

	return s.categoryRepo.Create(ctx, category)
}

func (s *CategoryService) List(ctx context.Context) ([]domain.Category, error) {
	return s.categoryRepo.List(ctx)
}
