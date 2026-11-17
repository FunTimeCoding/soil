package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreatePowerCable(
	x context.Context,
	outletDevice string,
	outlet string,
	portDevice string,
	port string,
) (client.Cable, error) {
	r, e := c.client.CreatePowerCableWithResponse(
		x,
		client.CreatePowerCableRequest{
			OutletDevice: outletDevice,
			Outlet:       outlet,
			PortDevice:   portDevice,
			Port:         port,
		},
	)

	if e != nil {
		return client.Cable{}, e
	}

	if r.JSON201 == nil {
		return client.Cable{}, refusal.New(
			"create power cable",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
