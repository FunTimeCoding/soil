package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Close() {
	if c.sftp != nil {
		errors.LogClose(c.sftp)
	}

	errors.LogClose(c.client)
}
