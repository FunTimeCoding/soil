package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateVirtualMachine(
	_ context.Context,
	name string,
	body client.UpdateVirtualMachineRequest,
) (client.VirtualMachine, error) {
	if c.Failure != nil {
		return client.VirtualMachine{}, c.Failure
	}

	if c.UpdatedMachines == nil {
		c.UpdatedMachines = map[string]client.UpdateVirtualMachineRequest{}
	}

	c.UpdatedMachines[name] = body

	return client.VirtualMachine{}, nil
}
