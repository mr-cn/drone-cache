package backend

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/meltwater/drone-cache/cache"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
)

// ftpBackend is an S3 implementation of the Backend
type ftpBackend struct {
	client *sftp.Client
}

// newFTP returns a new FTP remote Backend implemented
func newFTP(client *sftp.Client) cache.Backend {
	return &ftpBackend{client}
}

// Get returns an io.Reader for reading the contents of the file
func (c *ftpBackend) Get(p string) (io.ReadCloser, error) {
	if _, err := c.client.Stat(p); err != nil {
		return nil, errors.Wrap(err, "could not get the object")
	}
	return c.client.Open(p)
}

// Put uploads the contents of the io.ReadSeeker
func (c *ftpBackend) Put(p string, src io.ReadSeeker) error {
	if err := c.ensureDir(p); err != nil {
		return errors.Wrap(err, "could not put the object")
	}

	dst, err := c.client.Create(p)
	if err != nil {
		return errors.Wrap(err, "could not crate the object in remote server")
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return errors.Wrap(err, "could not put the object")
}

func (c *ftpBackend) ensureDir(p string) error {
	dir := filepath.Dir(p)
	if _, err := c.client.Stat(dir); err != nil {
		if err := c.client.MkdirAll(dir); err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not create remote directory <%s>", dir))
		}
	}

	return nil
}
