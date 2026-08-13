package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualDisk(
	machine string,
	name string,
	size int32,
) (*virtual_disk.Disk, error) {
	m, e := c.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	q := netbox.NewVirtualDiskRequest(
		netbox.PatchedVirtualDiskRequestVirtualMachine{Int32: &m.Identifier},
		name,
		size,
	)
	result, _, f := c.client.VirtualizationAPI.VirtualizationVirtualDisksCreate(
		c.context,
	).VirtualDiskRequest(*q).Execute()

	if f != nil {
		return nil, f
	}

	return virtual_disk.New(result), nil
}
