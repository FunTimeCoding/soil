package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualInterface(
	vm *virtual_machine.Machine,
	name string,
) (*netbox.VMInterface, error) {
	q := netbox.NewWritableVMInterfaceRequest(
		netbox.BriefVirtualMachineRequestAsPatchedVirtualDiskRequestVirtualMachine(
			netbox.NewBriefVirtualMachineRequest(vm.Name),
		),
		name,
	)
	result, _, e := c.client.VirtualizationAPI.VirtualizationInterfacesCreate(
		c.context,
	).WritableVMInterfaceRequest(*q).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return result, nil
}
