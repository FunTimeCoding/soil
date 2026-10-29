package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListVirtualDisks(_ context.Context, _ string) ([]*client.VirtualDisk, error) {
	return nil, c.Failure
}
