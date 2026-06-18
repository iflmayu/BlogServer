package storage

import "io"

type Uploader interface {
	Upload(relPath string, reader io.Reader, size int64) (string, error)
}
