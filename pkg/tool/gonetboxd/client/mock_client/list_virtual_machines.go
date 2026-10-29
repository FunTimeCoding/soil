package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListVirtualMachines(_ context.Context) ([]*client.VirtualMachine, error) {
	return c.VirtualMachineList, c.Failure
}
