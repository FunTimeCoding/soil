package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func (c *Client) ReadFile(path string) []byte {
	f, e := c.sftpClient().Open(path)
	errors.PanicOnError(e)
	defer errors.PanicClose(f)
	b, e := io.ReadAll(f)
	errors.PanicOnError(e)

	return b
}
