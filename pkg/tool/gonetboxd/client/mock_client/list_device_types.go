package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListDeviceTypes(_ context.Context) ([]*client.DeviceType, error) {
	return nil, c.Failure
}
