package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
)

func (c *Client) ListDirectory(path string) []fs.FileInfo {
	entries, e := c.sftpClient().ReadDir(path)
	errors.PanicOnError(e)

	return entries
}
