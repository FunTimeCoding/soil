package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListManufacturers() string {
	result, e := c.client.ListManufacturers(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
