package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreatePrefix(_ context.Context, _ string, _ string, _ string) (client.Prefix, error) {
	return client.Prefix{}, c.Failure
}
