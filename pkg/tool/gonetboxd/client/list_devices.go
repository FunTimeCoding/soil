package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListDevices(query *string) string {
	result, e := c.client.ListDevices(
		c.context,
		&client.ListDevicesParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
