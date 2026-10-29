package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListTags(_ context.Context) ([]*client.Tag, error) {
	return nil, c.Failure
}
