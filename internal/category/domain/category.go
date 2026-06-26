package domain

import "BlogServer/internal/common/domain"

type Category struct {
	domain.BaseModel
	Name        string `gorm:"size:32;uniqueIndex;not null" json:"name"`
	Slug        string `gorm:"size:32;uniqueIndex;not null" json:"slug"`
	Description string `gorm:"size:256" json:"description"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"`
}
