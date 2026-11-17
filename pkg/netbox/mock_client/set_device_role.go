package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceRole(
	name string,
	roleName string,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	r, e := c.DeviceRole(roleName)

	if e != nil {
		return nil, e
	}

	d, f := c.changeDevice("SetDeviceRole", name, roleName)

	if f != nil {
		return nil, f
	}

	d.Role = r.Name

	return d, nil
}
