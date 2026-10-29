package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListDevices(_ context.Context, _ *string) ([]*client.Device, error) {
	return c.DeviceList, c.Failure
}
