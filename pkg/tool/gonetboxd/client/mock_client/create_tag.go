package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateTag(_ context.Context, _ string) (client.Tag, error) {
	return client.Tag{}, c.Failure
}
