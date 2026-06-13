package storage

import "io"

type Uploader interface {
	Upload(path string, reader io.Reader, size int64) (string, error)
}
