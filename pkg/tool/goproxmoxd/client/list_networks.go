package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListNetworks(
	x context.Context,
	node string,
	p *client.ListNetworksParams,
) ([]client.Network, error) {
	r, e := c.client.ListNetworksWithResponse(x, node, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list networks", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
