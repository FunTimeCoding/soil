package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListClusters() string {
	result, e := c.client.ListClusters(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
