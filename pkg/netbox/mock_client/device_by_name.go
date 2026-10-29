package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) DeviceByName(n string) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, d := range c.devices {
		if d.Name == n {
			return d, nil
		}
	}

	return nil, not_found.New(constant.DeviceKind, n)
}
