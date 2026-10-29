package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SetInterfacePhysicalAddress(
	x context.Context,
	device string,
	name string,
	address string,
) (client.PhysicalAddress, error) {
	r, e := c.client.SetInterfacePhysicalAddressWithResponse(
		x,
		device,
		name,
		client.PhysicalAddressRequest{Address: address},
	)

	if e != nil {
		return client.PhysicalAddress{}, e
	}

	if r.JSON200 == nil {
		return client.PhysicalAddress{}, refusal.New(
			"set interface physical address",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
