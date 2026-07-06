package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListPrefixes() string {
	result, e := c.client.ListPrefixes(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
