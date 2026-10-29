package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SetVirtualInterfacePhysicalAddress(
	x context.Context,
	machine string,
	name string,
	address string,
) (client.PhysicalAddress, error) {
	r, e := c.client.SetVirtualInterfacePhysicalAddressWithResponse(
		x,
		machine,
		name,
		client.PhysicalAddressRequest{Address: address},
	)

	if e != nil {
		return client.PhysicalAddress{}, e
	}

	if r.JSON200 == nil {
		return client.PhysicalAddress{}, refusal.New(
			"set virtual interface physical address",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
