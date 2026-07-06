package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateTunnel(
	name string,
	encapsulation string,
	group string,
) string {
	result, e := c.client.CreateTunnel(
		c.context,
		client.CreateTunnelRequest{
			Name:          name,
			Encapsulation: encapsulation,
			Group:         group,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
