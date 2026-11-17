package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) SetInterfaceEthernetPower(
	d *device.Device,
	name string,
	mode string,
	kind string,
) (*network.Interface, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	i, e := c.DeviceInterface(d, name)

	if e != nil {
		return nil, e
	}

	i.EthernetPowerMode = mode
	i.EthernetPowerType = kind
	c.record("SetInterfaceEthernetPower", d.Name, name)

	return i, nil
}
