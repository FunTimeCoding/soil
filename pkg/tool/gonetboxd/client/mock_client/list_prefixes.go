package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListPrefixes(_ context.Context) ([]*client.Prefix, error) {
	return nil, c.Failure
}
