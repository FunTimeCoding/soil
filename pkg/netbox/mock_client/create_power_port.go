package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePowerPort(
	d *device.Device,
	name string,
) (*power_port.Port, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	v := netbox.NewPowerPortWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	v.Device.SetName(d.Name)
	result := power_port.New(v)
	c.ports[d.Identifier] = append(c.ports[d.Identifier], result)
	c.record("CreatePowerPort", d.Name, name)

	return result, nil
}
