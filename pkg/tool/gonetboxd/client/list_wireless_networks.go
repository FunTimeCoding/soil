package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListWirelessNetworks(
	x context.Context,
) ([]*client.WirelessNetwork, error) {
	r, e := c.client.ListWirelessNetworksWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list wireless networks",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
