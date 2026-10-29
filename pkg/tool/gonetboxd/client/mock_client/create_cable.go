package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateCable(_ context.Context, _ string, _ string, _ string, _ string) (client.Cable, error) {
	return client.Cable{}, c.Failure
}
