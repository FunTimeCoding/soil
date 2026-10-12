package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) changeDevice(
	method string,
	name string,
	value string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	c.record(method, name, value)

	return d, nil
}
