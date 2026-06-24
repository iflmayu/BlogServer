package service

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/article/repo"
	"context"
)

type ArticleService struct {
	articleRepo *repo.ArticleRepo
}

func NewArticleService(articleRepo *repo.ArticleRepo) *ArticleService {
	return &ArticleService{articleRepo: articleRepo}
}

type CreateArticleInput struct {
	Title      string
	Abstract   string
	Content    string
	Cover      string
	CategoryID uint
	Tags       []string
}

func (s *ArticleService) Create(ctx context.Context, input CreateArticleInput) error {
	article := &domain.Article{
		Title:      input.Title,
		Abstract:   input.Abstract,
		Content:    input.Content,
		Cover:      input.Cover,
		CategoryID: input.CategoryID,
		Tags:       domain.StringArray(input.Tags),
		Status:     domain.ArticleStatusPublished,
	}
	return s.articleRepo.Create(ctx, article)
}

type ListArticleInput struct {
	Page       int
	PageSize   int
	Keyword    string
	CategoryID uint
	Status     domain.ArticleStatus
}

func (s *ArticleService) List(ctx context.Context, input ListArticleInput) ([]domain.Article, int64, error) {
	return s.articleRepo.List(ctx, &repo.ListArticleQuery{
		Page:       input.Page,
		PageSize:   input.PageSize,
		Keyword:    input.Keyword,
		CategoryID: input.CategoryID,
		Status:     input.Status,
	})
}
