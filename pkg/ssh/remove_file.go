package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) RemoveFile(path string) {
	errors.PanicOnError(c.sftpClient().Remove(path))
}
