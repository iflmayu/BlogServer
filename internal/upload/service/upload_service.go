package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UploadService struct {
	uploadPath string
	baseURL    string
}

func NewUploadService(uploadPath, baseURL string) *UploadService {
	return &UploadService{
		uploadPath: uploadPath,
		baseURL:    baseURL,
	}
}

func (s *UploadService) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	// 校验文件类型
	ext := filepath.Ext(fileHeader.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" {
		return "", fmt.Errorf("仅支持 jpg/png/gif/webp 格式")
	}

	// 校验文件大小（比如 5MB）
	if fileHeader.Size > 5*1024*1024 {
		return "", fmt.Errorf("图片大小不能超过 5MB")
	}

	// 打开上传文件
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 生成目录：uploads/images/2026/06/
	now := time.Now()
	dir := filepath.Join(s.uploadPath, "images", now.Format("2006"), now.Format("01"))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// 生成文件名
	filename := uuid.New().String() + ext
	dstPath := filepath.Join(dir, filename)

	// 保存文件
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// 返回访问 URL
	relPath := filepath.Join("images", now.Format("2006"), now.Format("01"), filename)
	return s.baseURL + "/" + relPath, nil
}
