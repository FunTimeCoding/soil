package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func (c *Client) ReadFile(path string) ([]byte, error) {
	s, e := c.sftpClient()

	if e != nil {
		return nil, e
	}

	f, e := s.Open(path)

	if e != nil {
		return nil, c.refused("read", path, e)
	}

	defer errors.LogClose(f)
	b, e := io.ReadAll(f)

	if e != nil {
		return nil, c.refused("read", path, e)
	}

	return b, nil
}
