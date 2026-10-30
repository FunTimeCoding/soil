package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_disk"

func (c *Client) CreateVirtualDisk(
	machine string,
	name string,
	size int32,
) (*virtual_disk.Disk, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	m, e := c.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return c.SeedVirtualDisk(m, name, size, nil), nil
}
