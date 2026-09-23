package netbox

import (
	"github.com/funtimecoding/soil/pkg/integers64"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) UpdateVirtualInterface(
	vm *virtual_machine.Machine,
	name string,
	h net.HardwareAddr,
) (*netbox.VMInterface, error) {
	p, e := c.EnsurePhysical(h)

	if e != nil {
		return nil, e
	}

	i, f := c.VirtualMachineInterfaceByName(vm, name)

	if f != nil {
		return nil, f
	}

	assigned := p.ObjectType == constant.VirtualInterfaceAddress &&
		integers64.To32(p.ObjectIdentifier) == i.GetId()

	if !assigned {
		if _, g := c.AssignPhysicalToVirtualInterface(p, i.GetId()); g != nil {
			return nil, g
		}
	}

	q := netbox.NewWritableVMInterfaceRequest(
		netbox.BriefVirtualMachineRequestAsPatchedVirtualDiskRequestVirtualMachine(
			netbox.NewBriefVirtualMachineRequest(vm.Name),
		),
		name,
	)
	q.SetPrimaryMacAddress(
		netbox.BriefMACAddressRequestAsInterfaceRequestPrimaryMacAddress(
			netbox.NewBriefMACAddressRequest(h.String()),
		),
	)
	result, _, k := c.client.VirtualizationAPI.VirtualizationInterfacesUpdate(
		c.context,
		i.GetId(),
	).WritableVMInterfaceRequest(*q).Execute()

	if k != nil {
		return nil, k
	}

	return result, nil
}
