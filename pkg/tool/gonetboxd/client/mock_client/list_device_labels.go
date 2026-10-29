package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListDeviceLabels(_ context.Context, _ string) ([]*client.Label, error) {
	return nil, c.Failure
}
