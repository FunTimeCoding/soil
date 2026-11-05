package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) UpdateVirtualInterface(
	_ *virtual_machine.Machine,
	_ string,
	_ net.HardwareAddr,
) (*netbox.VMInterface, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
