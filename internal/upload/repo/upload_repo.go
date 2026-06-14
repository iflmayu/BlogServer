package repo

import (
	"context"

	"gorm.io/gorm"

	"BlogServer/internal/upload/domain"
)

type UploadRepo struct {
	db *gorm.DB
}

func NewUploadRepo(db *gorm.DB) *UploadRepo {
	return &UploadRepo{db: db}
}

// Create 创建上传记录
func (r *UploadRepo) Create(ctx context.Context, upload *domain.Upload) error {
	return r.db.WithContext(ctx).Create(upload).Error
}
