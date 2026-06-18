package domain

import "BlogServer/internal/common/domain"

type Upload struct {
	domain.BaseModel
	Filename string `gorm:"size:128" json:"filename"`
	URL      string `gorm:"size:512" json:"url"`
	Path     string `gorm:"size:512" json:"path"`
	Size     int64  `json:"size"`
	MimeType string `gorm:"size:64" json:"mime_type"`
}
