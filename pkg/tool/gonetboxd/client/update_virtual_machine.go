package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateVirtualMachine(
	x context.Context,
	name string,
	body client.UpdateVirtualMachineRequest,
) (client.VirtualMachine, error) {
	r, e := c.client.UpdateVirtualMachineWithResponse(x, name, body)

	if e != nil {
		return client.VirtualMachine{}, e
	}

	if r.JSON200 == nil {
		return client.VirtualMachine{}, refusal.New(
			"update virtual machine",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
