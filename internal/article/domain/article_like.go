package domain

import "BlogServer/internal/common/domain"

type ArticleLike struct {
	domain.BaseModel
	ArticleID uint `gorm:"index:idx_article_user,unique;not null" json:"article_id"`
	UserID    uint `gorm:"index:idx_article_user,unique;not null" json:"user_id"`
}
