package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) RenameDevice(
	name string,
	newName string,
) (*device.Device, error) {
	for i, d := range c.devices {
		if d.Name != name {
			continue
		}

		c.record("RenameDevice", name, newName)
		d.Raw.SetName(newName)
		c.devices[i] = device.New(d.Raw)

		return c.devices[i], nil
	}

	return nil, not_found.New(constant.DeviceKind, name)
}
