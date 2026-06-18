package storage

import (
	"io"
	"net/url"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	baseURL string
}

func NewLocalStorage(baseURL string) *LocalStorage {
	return &LocalStorage{baseURL: baseURL}
}

func (l *LocalStorage) Upload(relPath string, reader io.Reader, size int64) (string, error) {
	if err := os.MkdirAll(filepath.Dir(relPath), os.ModePerm); err != nil {
		return "", err
	}

	file, err := os.Create(relPath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	if _, err := io.Copy(file, reader); err != nil {
		return "", err
	}
	urlPath, _ := url.JoinPath(l.baseURL, "api", relPath)
	return urlPath, nil
}
