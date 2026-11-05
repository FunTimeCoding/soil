package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineInterfaces(_ *virtual_machine.Machine) ([]*network.Interface, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
