package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetVirtualMachine(
	x context.Context,
	name string,
) (client.VirtualMachine, error) {
	r, e := c.client.GetVirtualMachineWithResponse(x, name)

	if e != nil {
		return client.VirtualMachine{}, e
	}

	if r.JSON200 == nil {
		return client.VirtualMachine{}, refusal.New(
			"get virtual machine",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
