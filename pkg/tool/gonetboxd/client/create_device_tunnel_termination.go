package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateDeviceTunnelTermination(
	x context.Context,
	device string,
	tunnel string,
	interfaceName string,
	role string,
) (client.TunnelTermination, error) {
	r, e := c.client.CreateDeviceTunnelTerminationWithResponse(
		x,
		device,
		client.CreateTunnelTerminationRequest{
			Tunnel:    tunnel,
			Interface: interfaceName,
			Role:      role,
		},
	)

	if e != nil {
		return client.TunnelTermination{}, e
	}

	if r.JSON201 == nil {
		return client.TunnelTermination{}, refusal.New(
			"create device tunnel termination",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
