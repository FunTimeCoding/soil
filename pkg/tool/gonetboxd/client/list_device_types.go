package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListDeviceTypes() string {
	result, e := c.client.ListDeviceTypes(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
