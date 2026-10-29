package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateDeviceRole(_ context.Context, _ string, _ bool) (client.DeviceRole, error) {
	return client.DeviceRole{}, c.Failure
}
