package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) SetInterfacePhysicalAddress(_ context.Context, _ string, _ string, _ string) (client.PhysicalAddress, error) {
	return client.PhysicalAddress{}, c.Failure
}
