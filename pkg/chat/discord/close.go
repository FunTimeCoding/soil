package discord

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Close() {
	errors.LogClose(c.client)
}
