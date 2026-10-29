package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateVirtualMachine(
	x context.Context,
	name string,
	cluster string,
) (client.VirtualMachine, error) {
	r, e := c.client.CreateVirtualMachineWithResponse(
		x,
		client.CreateVirtualMachineRequest{Name: name, Cluster: cluster},
	)

	if e != nil {
		return client.VirtualMachine{}, e
	}

	if r.JSON201 == nil {
		return client.VirtualMachine{}, refusal.New(
			"create virtual machine",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
