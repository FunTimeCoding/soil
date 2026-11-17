package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceLocalContext(
	name string,
	payload map[string]any,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	d, e := c.Device(name)

	if e != nil {
		return nil, e
	}

	d.LocalContext = nil

	if len(payload) > 0 {
		d.LocalContext = payload
	}

	c.record("SetDeviceLocalContext", name)

	return d, nil
}
