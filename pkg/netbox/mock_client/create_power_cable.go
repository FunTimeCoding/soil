package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePowerCable(
	o *power_outlet.Outlet,
	p *power_port.Port,
) (*cable.Cable, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	o.ConnectedDevice = p.Device
	o.ConnectedPort = p.Name
	c.record("CreatePowerCable", o.Name, p.Name)
	v := netbox.NewCableWithDefaults()
	v.SetId(c.nextIdentifier())

	return cable.New(v), nil
}
