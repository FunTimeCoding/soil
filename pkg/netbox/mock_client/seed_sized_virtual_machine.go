package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedSizedVirtualMachine(
	name string,
	cores float64,
	memory int32,
	place string,
) *virtual_machine.Machine {
	v := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	v.SetVcpus(cores)
	v.SetMemory(memory)
	v.SetSite(*netbox.NewBriefSite(0, "", place, place, place))
	result := virtual_machine.New(v)
	c.machines = append(c.machines, result)

	return result
}
