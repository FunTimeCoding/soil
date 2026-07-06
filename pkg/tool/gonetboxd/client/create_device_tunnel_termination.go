package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateDeviceTunnelTermination(
	device string,
	tunnel string,
	interfaceName string,
	role string,
) string {
	result, e := c.client.CreateDeviceTunnelTermination(
		c.context,
		device,
		client.CreateTunnelTerminationRequest{
			Tunnel:    tunnel,
			Interface: interfaceName,
			Role:      role,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
