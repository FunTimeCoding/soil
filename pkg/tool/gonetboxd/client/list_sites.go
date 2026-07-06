package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListSites() string {
	result, e := c.client.ListSites(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
