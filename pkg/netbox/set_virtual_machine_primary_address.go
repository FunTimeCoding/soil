package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachinePrimaryAddress(
	name string,
	address string,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	addresses, f := c.VirtualMachineAddresses(name)

	if f != nil {
		return nil, f
	}

	match := internet_address.Find(addresses, address)

	if match == nil {
		return nil, fmt.Errorf(
			"address %s not assigned to virtual machine %s",
			address,
			name,
		)
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetPrimaryIp4(
		netbox.DeviceWithConfigContextRequestPrimaryIp4{
			Int32: &match.Identifier,
		},
	)

	return c.updateVirtualMachine(m, q)
}
