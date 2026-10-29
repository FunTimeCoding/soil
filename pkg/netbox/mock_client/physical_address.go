package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddress(_ net.HardwareAddr) (*physical_address.Address, error) {
	return nil, c.Failure
}
