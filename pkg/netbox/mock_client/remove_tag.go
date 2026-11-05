package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) RemoveTag(
	_ string,
	_ string,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
