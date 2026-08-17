package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Log(limit *int) string {
	result, e := c.client.FirewallLog(
		c.context,
		&client.FirewallLogParams{Limit: limit},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
