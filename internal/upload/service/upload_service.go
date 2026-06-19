package service

import (
	"BlogServer/internal/upload/domain"
	"BlogServer/internal/upload/repo"
	"BlogServer/pkg/config"
	"BlogServer/pkg/storage"
	"context"
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type UploadService struct {
	repo    *repo.UploadRepo
	storage storage.Uploader
	cfg     config.Upload
}

func NewUploadService(repo *repo.UploadRepo, storage storage.Uploader, cfg config.Upload) *UploadService {
	return &UploadService{
		repo:    repo,
		storage: storage,
		cfg:     cfg,
	}
}

func (s *UploadService) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	// 校验文件类型
	ext := filepath.Ext(fileHeader.Filename)
	if !slices.Contains(s.cfg.AllowedTypes, ext) {
		return "", fmt.Errorf("仅支持 %s 格式", strings.Join(s.cfg.AllowedTypes, " "))
	}

	// 校验文件大小
	if fileHeader.Size > int64(s.cfg.MaxSize)*1024*1024 {
		return "", fmt.Errorf("图片大小不能超过 %dMB", s.cfg.MaxSize)
	}
	// 打开上传文件
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func() {
		_ = src.Close()
	}()

	// 生成文件名, 相对路径
	filename := uuid.New().String() + ext
	relPath := path.Join(s.cfg.UploadDir, filename)

	// 调用 Storage 层保存（传入 relPath、reader、size）
	url, err := s.storage.Upload(relPath, src, fileHeader.Size)
	if err != nil {
		return "", err
	}

	// 保存上传记录到数据库
	upload := &domain.Upload{
		Filename: filename,
		URL:      url,
		Path:     relPath,
		Size:     fileHeader.Size,
		MimeType: fileHeader.Header.Get("Content-Type"),
	}
	if err := s.repo.Create(ctx, upload); err != nil {
		return "", err
	}

	return url, nil
}
