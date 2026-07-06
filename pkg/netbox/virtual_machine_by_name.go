package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineByName(n string) (*virtual_machine.Machine, error) {
	result, e := c.VirtualMachines()

	if e != nil {
		return nil, e
	}

	for _, m := range result {
		if m.Name == n {
			return m, nil
		}
	}

	return nil, fmt.Errorf("virtual machine not found: %s", n)
}
