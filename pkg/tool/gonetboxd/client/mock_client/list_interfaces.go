package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListInterfaces(_ context.Context, _ string) ([]*client.Interface, error) {
	return nil, c.Failure
}
