package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePowerOutlet(
	d *device.Device,
	name string,
) (*power_outlet.Outlet, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	v := netbox.NewPowerOutletWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	v.Device.SetName(d.Name)
	result := power_outlet.New(v)
	c.outlets[d.Identifier] = append(c.outlets[d.Identifier], result)
	c.record("CreatePowerOutlet", d.Name, name)

	return result, nil
}
