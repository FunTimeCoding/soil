package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListTunnels(x context.Context) ([]*client.Tunnel, error) {
	r, e := c.client.ListTunnelsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list tunnels", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
