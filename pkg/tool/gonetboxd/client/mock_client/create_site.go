package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateSite(_ context.Context, _ string) (client.Site, error) {
	return client.Site{}, c.Failure
}
