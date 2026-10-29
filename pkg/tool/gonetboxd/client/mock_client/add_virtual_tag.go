package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) AddVirtualTag(_ context.Context, _ string, _ string) (client.VirtualMachine, error) {
	return client.VirtualMachine{}, c.Failure
}
