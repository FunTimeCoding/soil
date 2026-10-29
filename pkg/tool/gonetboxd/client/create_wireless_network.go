package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateWirelessNetwork(
	x context.Context,
	ssid string,
) (client.WirelessNetwork, error) {
	r, e := c.client.CreateWirelessNetworkWithResponse(
		x,
		client.CreateNameRequest{Name: ssid},
	)

	if e != nil {
		return client.WirelessNetwork{}, e
	}

	if r.JSON201 == nil {
		return client.WirelessNetwork{}, refusal.New(
			"create wireless network",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
