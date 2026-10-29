package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineDisks(_ *virtual_machine.Machine) ([]*virtual_disk.Disk, error) {
	return nil, nil
}
