package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListInstances() string {
	result, e := c.client.ListInstancesWithResponse(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
