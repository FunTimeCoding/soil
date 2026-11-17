package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceSite(
	name string,
	siteName string,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	s, e := c.Site(siteName)

	if e != nil {
		return nil, e
	}

	d, f := c.changeDevice("SetDeviceSite", name, siteName)

	if f != nil {
		return nil, f
	}

	d.Site = s.Name

	return d, nil
}
