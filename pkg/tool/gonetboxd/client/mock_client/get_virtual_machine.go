package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) GetVirtualMachine(_ context.Context, _ string) (client.VirtualMachine, error) {
	return client.VirtualMachine{}, c.Failure
}
