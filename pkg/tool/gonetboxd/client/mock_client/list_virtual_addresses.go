package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListVirtualAddresses(_ context.Context, _ string) ([]*client.Address, error) {
	return nil, c.Failure
}
