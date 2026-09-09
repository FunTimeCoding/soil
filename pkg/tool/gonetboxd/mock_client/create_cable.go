package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) CreateCable(
	_ *network.Interface,
	_ *network.Interface,
) (*cable.Cable, error) {
	return nil, nil
}
