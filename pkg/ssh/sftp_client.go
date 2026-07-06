package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/pkg/sftp"
)

func (c *Client) sftpClient() *sftp.Client {
	if c.sftp != nil {
		return c.sftp
	}

	result, e := sftp.NewClient(c.client)
	errors.PanicOnError(e)
	c.sftp = result

	return c.sftp
}
