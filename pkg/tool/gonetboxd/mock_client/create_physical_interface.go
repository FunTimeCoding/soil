package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) CreatePhysicalInterface(
	_ net.HardwareAddr,
	_ string,
	_ *network.Interface,
) (*physical_address.Address, error) {
	return nil, nil
}
