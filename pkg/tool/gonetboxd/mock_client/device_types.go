package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device_type"

func (c *Client) DeviceTypes() ([]*device_type.Type, error) {
	return nil, nil
}
