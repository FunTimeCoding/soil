package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedVirtualMachine(name string) *virtual_machine.Machine {
	v := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	result := virtual_machine.New(v)
	c.machines = append(c.machines, result)

	return result
}
