package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Close() {
	if c.sftp != nil {
		errors.LogClose(c.sftp)
	}

	if c.client != nil {
		errors.LogClose(c.client)
	}
}
