package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListSites(_ context.Context) ([]*client.Site, error) {
	return nil, c.Failure
}
