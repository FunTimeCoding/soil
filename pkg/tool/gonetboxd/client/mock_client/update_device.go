package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateDevice(_ context.Context, _ string, _ client.UpdateDeviceRequest) (client.Device, error) {
	return client.Device{}, c.Failure
}
