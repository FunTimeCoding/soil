package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateVirtualInterface(_ context.Context, _ string, _ string) (client.VirtualInterface, error) {
	return client.VirtualInterface{}, c.Failure
}
