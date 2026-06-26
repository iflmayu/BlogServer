package service

import (
	aRepo "BlogServer/internal/article/repo"
	"BlogServer/internal/category/domain"
	"BlogServer/internal/category/repo"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepo *repo.CategoryRepo
	articleRepo  *aRepo.ArticleRepo
}

func NewCategoryService(
	categoryRepo *repo.CategoryRepo,
	articleRepo *aRepo.ArticleRepo,
) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
		articleRepo:  articleRepo}
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

type UpdateCategoryInput struct {
	ID          uint
	Name        string
	Slug        string
	Description string
	SortOrder   int
}

func (s *CategoryService) Update(ctx context.Context, input UpdateCategoryInput) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("分类名称不能为空")
	}
	if strings.TrimSpace(input.Slug) == "" {
		return errors.New("分类标识不能为空")
	}

	category, err := s.categoryRepo.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("分类不存在")
		}
		return err
	}

	existing, err := s.categoryRepo.GetByNameOrSlug(ctx, input.Name, input.Slug, input.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		if existing.Name == input.Name {
			return errors.New("分类名称已存在")
		}
		return errors.New("分类标识已存在")
	}

	category.Name = input.Name
	category.Slug = input.Slug
	category.Description = input.Description
	category.SortOrder = input.SortOrder

	return s.categoryRepo.Update(ctx, category)
}

func (s *CategoryService) Delete(ctx context.Context, id uint) error {
	category, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("分类不存在")
		}
		return err
	}

	count, err := s.articleRepo.CountByCategoryID(ctx, category.ID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该分类下存在文章，无法删除")
	}

	return s.categoryRepo.Delete(ctx, category.ID)
}
