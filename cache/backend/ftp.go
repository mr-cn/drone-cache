package backend

import (
	"io"
	"time"

	"github.com/meltwater/drone-cache/cache"
)

// ftpBacked is an S3 implementation of the Backend
type ftpBacked struct {
	user     string
	password string

	timeout            time.Duration
	connectionsPerHost int

	c FTPConfig
}

// newFTP returns a new FTP remote Backend implemented
func newFTP(c FTPConfig) cache.Backend {
	return &ftpBacked{c: c}
}

// Get returns an io.Reader for reading the contents of the file
func (c *ftpBacked) Get(p string) (io.ReadCloser, error) {
	return nil, nil
}

// Put uploads the contents of the io.ReadSeeker
func (c *ftpBacked) Put(p string, src io.ReadSeeker) error {
	return nil
}
