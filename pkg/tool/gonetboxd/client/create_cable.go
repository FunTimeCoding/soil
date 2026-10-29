package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateCable(
	x context.Context,
	deviceA string,
	interfaceA string,
	deviceB string,
	interfaceB string,
) (client.Cable, error) {
	r, e := c.client.CreateCableWithResponse(
		x,
		client.CreateCableRequest{
			DeviceA:    deviceA,
			InterfaceA: interfaceA,
			DeviceB:    deviceB,
			InterfaceB: interfaceB,
		},
	)

	if e != nil {
		return client.Cable{}, e
	}

	if r.JSON201 == nil {
		return client.Cable{}, refusal.New(
			"create cable",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
