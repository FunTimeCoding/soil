package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListDeviceRoles(_ context.Context) ([]*client.DeviceRole, error) {
	return nil, c.Failure
}
