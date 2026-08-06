package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachinePlatform(
	name string,
	platformName string,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	p, f := c.PlatformByName(platformName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetPlatform(
		netbox.DeviceTypeRequestDefaultPlatform{
			Int32: &p.Identifier,
		},
	)

	return c.updateVirtualMachine(m, q)
}
