package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) VirtualMachineInterfaceByName(
	vm *virtual_machine.Machine,
	name string,
) (*netbox.VMInterface, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationInterfacesList(
		c.context,
	).VirtualMachineId([]int32{vm.Identifier}).Execute()

	if e != nil {
		return nil, e
	}

	for _, i := range result.Results {
		if i.GetName() == name {
			return &i, nil
		}
	}

	return nil, fmt.Errorf(
		"interface %s not found for VM %s",
		name,
		vm.Name,
	)
}
