package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) WriteFile(
	path string,
	content []byte,
) error {
	s, e := c.sftpClient()

	if e != nil {
		return e
	}

	f, e := s.Create(path)

	if e != nil {
		return c.refused("write", path, e)
	}

	if _, g := f.Write(content); g != nil {
		errors.LogClose(f)

		return c.refused("write", path, g)
	}

	if h := f.Close(); h != nil {
		return c.refused("write", path, h)
	}

	return nil
}
