package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListAddresses(name string) string {
	result, e := c.client.ListAddresses(c.context, name)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
