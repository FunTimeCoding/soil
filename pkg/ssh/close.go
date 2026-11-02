package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func (c *Client) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.sftp != nil {
		if e := c.sftp.Close(); !errors.Is(e, io.EOF) {
			errors.LogOnError(e)
		}

		c.sftp = nil
	}

	if c.client != nil {
		errors.LogReclose(c.client)
		c.client = nil
	}
}
