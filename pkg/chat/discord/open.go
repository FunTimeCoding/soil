package discord

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Open() {
	errors.PanicOnError(c.client.Open())
}
