package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedVirtualDisk(
	m *virtual_machine.Machine,
	name string,
	size int32,
	fields map[string]any,
) *virtual_disk.Disk {
	v := netbox.NewVirtualDiskWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	v.SetSize(size)
	v.SetCustomFields(fields)
	v.SetVirtualMachine(
		*netbox.NewBriefVirtualMachine(m.Identifier, "", m.Name, m.Name),
	)
	result := virtual_disk.New(v)
	c.disks[m.Identifier] = append(c.disks[m.Identifier], result)

	return result
}
