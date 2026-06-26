package repo

import (
	"BlogServer/internal/article/domain"
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *ArticleRepo) GetByID(ctx context.Context, id uint) (*domain.Article, error) {
	var article domain.Article
	if err := r.db.WithContext(ctx).First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepo) Update(ctx context.Context, article *domain.Article) error {
	return r.db.WithContext(ctx).Model(&domain.Article{}).
		Where("id = ?", article.ID).
		Updates(map[string]interface{}{
			"title":       article.Title,
			"abstract":    article.Abstract,
			"content":     article.Content,
			"cover":       article.Cover,
			"category_id": article.CategoryID,
			"tags":        article.Tags,
			"status":      article.Status,
		}).Error
}

// ToggleLike 切换点赞状态，返回切换后的点赞状态和最新点赞数
func (r *ArticleRepo) ToggleLike(ctx context.Context, articleID, userID uint) (bool, int64, error) {
	var isLiked bool
	var likeCount int64

	// 开启数据库事务
	//如果 func(tx *gorm.DB) error 返回 nil，事务提交；返回非 nil，事务回滚。
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 锁住文章行，串行化该文章的所有点赞操作
		var article domain.Article
		// tx.Clauses(clause.Locking{Strength: "UPDATE"})：加 FOR UPDATE 行锁
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&article, articleID).Error; err != nil {
			return err
		}

		var like domain.ArticleLike
		err := tx.Where("article_id = ? AND user_id = ?", articleID, userID).First(&like).Error

		if err == nil {
			// 已点赞 -> 取消点赞
			if err := tx.Delete(&like).Error; err != nil {
				return err
			}
			article.LikeCount--
			isLiked = false
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未点赞 -> 点赞
			like = domain.ArticleLike{ArticleID: articleID, UserID: userID}
			if err := tx.Create(&like).Error; err != nil {
				return err
			}
			article.LikeCount++
			isLiked = true
		} else {
			return err
		}

		// 同步更新文章点赞数
		if err := tx.Model(&article).Update("like_count", article.LikeCount).Error; err != nil {
			return err
		}

		likeCount = article.LikeCount
		return nil
	})

	return isLiked, likeCount, err
}

// HasLiked 查询用户是否点赞过某篇文章
func (r *ArticleRepo) HasLiked(ctx context.Context, articleID, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.ArticleLike{}).
		Where("article_id = ? AND user_id = ?", articleID, userID).
		Count(&count).Error
	return count > 0, err
}

// IncrementViewCount 增加文章浏览量
func (r *ArticleRepo) IncrementViewCount(ctx context.Context, articleID uint) (int64, error) {
	result := r.db.WithContext(ctx).Model(&domain.Article{}).
		Where("id = ?", articleID).
		Update("view_count", gorm.Expr("view_count + 1"))
	if result.Error != nil {
		return 0, result.Error
	}
	// 判断文章 ID 是否存在
	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	var article domain.Article
	if err := r.db.WithContext(ctx).First(&article, articleID).Error; err != nil {
		return 0, err
	}
	return article.ViewCount, nil
}

// Delete 删除文章及其关联的点赞记录
func (r *ArticleRepo) Delete(ctx context.Context, articleID uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("article_id = ?", articleID).Delete(&domain.ArticleLike{}).Error; err != nil {
			return err
		}

		result := tx.Delete(&domain.Article{}, articleID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})
}

func (r *ArticleRepo) CountByCategoryID(ctx context.Context, categoryID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.Article{}).
		Where("category_id = ?", categoryID).
		Count(&count).Error
	return count, err
}
