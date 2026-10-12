package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) CreateCable(
	a *network.Interface,
	b *network.Interface,
) (*cable.Cable, error) {
	c.record("CreateCable", a.Name, b.Name)

	return nil, nil
}
