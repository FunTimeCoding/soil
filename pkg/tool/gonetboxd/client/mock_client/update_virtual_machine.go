package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateVirtualMachine(_ context.Context, _ string, _ client.UpdateVirtualMachineRequest) (client.VirtualMachine, error) {
	return client.VirtualMachine{}, c.Failure
}
