package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedInterface(
	d *device.Device,
	name string,
) *network.Interface {
	v := netbox.NewInterfaceWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	result := network.New(v)
	c.interfaces[d.Identifier] = append(c.interfaces[d.Identifier], result)

	return result
}
