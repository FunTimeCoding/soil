package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateDeviceType(_ context.Context, _ string, _ string) (client.DeviceType, error) {
	return client.DeviceType{}, c.Failure
}
