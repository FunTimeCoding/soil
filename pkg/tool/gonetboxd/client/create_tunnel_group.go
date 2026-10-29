package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateTunnelGroup(
	x context.Context,
	name string,
) (client.TunnelGroup, error) {
	r, e := c.client.CreateTunnelGroupWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.TunnelGroup{}, e
	}

	if r.JSON201 == nil {
		return client.TunnelGroup{}, refusal.New(
			"create tunnel group",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
