package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"maps"
)

func (c *Client) SetDeviceCustomFields(
	name string,
	fields map[string]any,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	d, e := c.Device(name)

	if e != nil {
		return nil, e
	}

	if d.CustomFields == nil {
		d.CustomFields = map[string]any{}
	}

	maps.Copy(d.CustomFields, fields)
	c.record("SetDeviceCustomFields", name)

	return d, nil
}
