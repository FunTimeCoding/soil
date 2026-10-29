package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateVirtualInterface(
	x context.Context,
	vmName string,
	name string,
) (client.VirtualInterface, error) {
	r, e := c.client.CreateVirtualInterfaceWithResponse(
		x,
		vmName,
		client.CreateVirtualInterfaceRequest{Name: name},
	)

	if e != nil {
		return client.VirtualInterface{}, e
	}

	if r.JSON201 == nil {
		return client.VirtualInterface{}, refusal.New(
			"create virtual interface",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
