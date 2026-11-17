package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceType(
	name string,
	model string,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	t, e := c.DeviceType(model)

	if e != nil {
		return nil, e
	}

	d, f := c.changeDevice("SetDeviceType", name, model)

	if f != nil {
		return nil, f
	}

	d.Type = t.Model

	return d, nil
}
