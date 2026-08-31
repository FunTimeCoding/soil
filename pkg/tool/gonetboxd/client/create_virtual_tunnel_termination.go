package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateVirtualTunnelTermination(
	vm string,
	tunnel string,
	interfaceName string,
	role string,
) *response.Response {
	result, e := c.client.CreateVirtualTunnelTermination(
		c.context,
		vm,
		client.CreateTunnelTerminationRequest{
			Tunnel:    tunnel,
			Interface: interfaceName,
			Role:      role,
		},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
