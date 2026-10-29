package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) RemoveDeviceTag(_ context.Context, _ string, _ string) (client.Device, error) {
	return client.Device{}, c.Failure
}
