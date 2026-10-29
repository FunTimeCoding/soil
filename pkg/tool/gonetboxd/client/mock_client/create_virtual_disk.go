package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateVirtualDisk(_ context.Context, _ string, _ string, _ int32) (client.VirtualDisk, error) {
	return client.VirtualDisk{}, c.Failure
}
