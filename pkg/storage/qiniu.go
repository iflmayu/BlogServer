package storage

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
)

type QiniuStorage struct {
	bucket string
	domain string
	client *uploader.UploadManager
}

func NewQiniuStorage(accessKey, secretKey, bucket, domain string) *QiniuStorage {
	mac := credentials.NewCredentials(accessKey, secretKey)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})

	return &QiniuStorage{
		bucket: bucket,
		domain: domain,
		client: uploadManager,
	}
}

func (q *QiniuStorage) Upload(relPath string, reader io.Reader, size int64) (string, error) {
	filename := path.Base(relPath)
	fmt.Println(filename)

	err := q.client.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: q.bucket,
		ObjectName: &relPath,
		FileName:   filename,
	}, nil)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", q.domain, relPath), nil
}
