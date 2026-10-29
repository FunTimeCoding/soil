package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"

func (c *Client) VirtualDeviceContexts() ([]*virtual_device_context.Context, error) {
	return nil, c.Failure
}
