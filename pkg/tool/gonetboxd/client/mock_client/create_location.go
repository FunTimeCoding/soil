package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateLocation(_ context.Context, _ string, _ string) (client.Location, error) {
	return client.Location{}, c.Failure
}
