package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) MustVirtualMachineInterfaceByName(
	vm *virtual_machine.Machine,
	name string,
) *netbox.VMInterface {
	result, e := c.VirtualMachineInterfaceByName(vm, name)
	errors.PanicOnError(e)

	return result
}
