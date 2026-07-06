package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) WriteFile(
	path string,
	content []byte,
) {
	f, e := c.sftpClient().Create(path)
	errors.PanicOnError(e)
	defer errors.PanicClose(f)
	_, e = f.Write(content)
	errors.PanicOnError(e)
}
