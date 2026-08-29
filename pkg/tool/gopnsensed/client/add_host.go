package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddHost(
	body client.HostRequest,
	apply *bool,
) string {
	result, e := c.client.AddHost(
		c.context,
		&client.AddHostParams{Apply: apply},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
