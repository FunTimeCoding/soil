package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateInterface(_ context.Context, _ string, _ string, _ string, _ string) (client.Interface, error) {
	return client.Interface{}, c.Failure
}
