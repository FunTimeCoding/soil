package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineDisks(
	m *virtual_machine.Machine,
) ([]*virtual_disk.Disk, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return c.disks[m.Identifier], nil
}
