package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListLocations(_ context.Context) ([]*client.Location, error) {
	return nil, c.Failure
}
