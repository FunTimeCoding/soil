package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Gear() string {
	result, e := c.client.GetGear(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
