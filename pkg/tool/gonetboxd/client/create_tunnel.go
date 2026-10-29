package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateTunnel(
	x context.Context,
	name string,
	encapsulation string,
	group string,
) (client.Tunnel, error) {
	r, e := c.client.CreateTunnelWithResponse(
		x,
		client.CreateTunnelRequest{
			Name:          name,
			Encapsulation: encapsulation,
			Group:         group,
		},
	)

	if e != nil {
		return client.Tunnel{}, e
	}

	if r.JSON201 == nil {
		return client.Tunnel{}, refusal.New(
			"create tunnel",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
