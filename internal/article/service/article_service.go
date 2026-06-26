package service

import (
	"BlogServer/internal/article/domain"
	"BlogServer/internal/article/repo"
	catRepo "BlogServer/internal/category/repo"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ArticleService struct {
	articleRepo  *repo.ArticleRepo
	categoryRepo *catRepo.CategoryRepo
}

func NewArticleService(
	articleRepo *repo.ArticleRepo,
	categoryRepo *catRepo.CategoryRepo,
) *ArticleService {
	return &ArticleService{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
	}
}

type CreateArticleInput struct {
	Title      string
	Abstract   string
	Content    string
	Cover      string
	CategoryID uint
	Tags       domain.StringArray
}

func (s *ArticleService) Create(ctx context.Context, input CreateArticleInput) error {
	if err := s.validateCategory(ctx, input.CategoryID); err != nil {
		return err
	}

	article := &domain.Article{
		Title:      input.Title,
		Abstract:   input.Abstract,
		Content:    input.Content,
		Cover:      input.Cover,
		CategoryID: input.CategoryID,
		Tags:       input.Tags,
		Status:     domain.ArticleStatusPublished,
	}
	return s.articleRepo.Create(ctx, article)
}

func (s *ArticleService) validateCategory(ctx context.Context, categoryID uint) error {
	if categoryID == 0 {
		return nil
	}
	_, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("分类不存在")
		}
		return err
	}
	return nil
}

type ListArticleInput struct {
	Page       int
	PageSize   int
	Keyword    string
	CategoryID uint
	Status     domain.ArticleStatus
}

type ListArticleItem struct {
	domain.Article
	CategoryName string `json:"category_name"`
}

func (s *ArticleService) List(ctx context.Context, input ListArticleInput) ([]ListArticleItem, int64, error) {
	repoItems, total, err := s.articleRepo.List(ctx, &repo.ListArticleQuery{
		Page:       input.Page,
		PageSize:   input.PageSize,
		Keyword:    input.Keyword,
		CategoryID: input.CategoryID,
		Status:     input.Status,
	})
	if err != nil {
		return nil, 0, err
	}

	items := make([]ListArticleItem, len(repoItems))
	for i, item := range repoItems {
		items[i] = ListArticleItem{
			Article:      item.Article,
			CategoryName: item.CategoryName,
		}
	}

	return items, total, nil
}

type UpdateArticleInput struct {
	ID         uint
	Title      string
	Abstract   string
	Content    string
	Cover      string
	CategoryID uint
	Tags       domain.StringArray
	Status     domain.ArticleStatus
}

// wrapNotFound 将 gorm 的 record not found 转换为友好提示
func wrapNotFound(err error, msg string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(msg)
	}
	return err
}

func (s *ArticleService) Update(ctx context.Context, input UpdateArticleInput) error {
	if !input.Status.IsValid() {
		return errors.New("无效的文章状态")
	}

	if err := s.validateCategory(ctx, input.CategoryID); err != nil {
		return err
	}

	article, err := s.articleRepo.GetByID(ctx, input.ID)
	if err != nil {
		return wrapNotFound(err, "文章不存在")
	}

	article.Title = input.Title
	article.Abstract = input.Abstract
	article.Content = input.Content
	article.Cover = input.Cover
	article.CategoryID = input.CategoryID
	article.Tags = input.Tags
	article.Status = input.Status

	return s.articleRepo.Update(ctx, article)
}

type ArticleDetail struct {
	domain.Article
	CategoryName string
}

func (s *ArticleService) GetArticleDetail(ctx context.Context, id uint) (*ArticleDetail, error) {
	article, err := s.articleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, wrapNotFound(err, "文章不存在")
	}
	if article.Status != domain.ArticleStatusPublished {
		return nil, errors.New("文章不存在或已下线")
	}

	detail := &ArticleDetail{Article: *article}
	if article.CategoryID > 0 {
		category, err := s.categoryRepo.GetByID(ctx, article.CategoryID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if category != nil {
			detail.CategoryName = category.Name
		}
	}

	return detail, nil
}

func (s *ArticleService) ToggleLike(ctx context.Context, articleID, userID uint) (bool, int64, error) {
	return s.articleRepo.ToggleLike(ctx, articleID, userID)
}

func (s *ArticleService) HasLiked(ctx context.Context, articleID, userID uint) (bool, error) {
	return s.articleRepo.HasLiked(ctx, articleID, userID)
}

func (s *ArticleService) ViewArticle(ctx context.Context, articleID uint) (int64, error) {
	return s.articleRepo.IncrementViewCount(ctx, articleID)
}

func (s *ArticleService) Delete(ctx context.Context, articleID uint) error {
	if err := s.articleRepo.Delete(ctx, articleID); err != nil {
		return wrapNotFound(err, "文章不存在")
	}
	return nil
}
