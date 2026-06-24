package domain

import "BlogServer/internal/common/domain"

type Article struct {
	domain.BaseModel
	Title        string        `gorm:"size:256;not null" json:"title"`
	Abstract     string        `gorm:"size:512" json:"abstract"`
	Content      string        `gorm:"type:text" json:"content"`
	Cover        string        `gorm:"size:512" json:"cover"`
	CategoryID   uint          `gorm:"index" json:"category_id"`
	Tags         StringArray   `gorm:"type:json" json:"tags"`
	Status       ArticleStatus `gorm:"default:1;not null;index" json:"status"`
	ViewCount    int64         `gorm:"default:0" json:"view_count"`
	LikeCount    int64         `gorm:"default:0" json:"like_count"`
	CommentCount int64         `gorm:"default:0" json:"comment_count"`
}
