package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineByName(n string) (*virtual_machine.Machine, error) {
	for _, m := range c.machines {
		if m.Name == n {
			return m, nil
		}
	}

	return nil, not_found.New(constant.VirtualMachineKind, n)
}
