package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	rootPath string
	baseURL  string
}

func NewLocalStorage(rootPath, baseURL string) *LocalStorage {
	return &LocalStorage{rootPath: rootPath, baseURL: baseURL}
}

func (l *LocalStorage) Upload(path string, reader io.Reader, size int64) (string, error) {
	fullPath := filepath.Join(l.rootPath, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return "", err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	if _, err := io.Copy(file, reader); err != nil {
		return "", err
	}

	return l.baseURL + "/" + path, nil
}
