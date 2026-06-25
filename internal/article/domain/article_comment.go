package domain

import "BlogServer/internal/common/domain"

type ArticleComment struct {
	domain.BaseModel
	ArticleID uint          `gorm:"index;not null" json:"article_id"`
	UserID    uint          `gorm:"index;not null" json:"user_id"`
	AtID      uint          `gorm:"index;default:0" json:"at_id"`
	Content   string        `gorm:"type:text;not null" json:"content"`
	Status    CommentStatus `gorm:"default:1;not null;index" json:"status"`
}
