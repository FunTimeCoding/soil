package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateCable(
	a *network.Interface,
	b *network.Interface,
) (*cable.Cable, error) {
	c.record("CreateCable", a.Name, b.Name)
	v := netbox.NewCableWithDefaults()
	v.SetId(c.nextIdentifier())

	return cable.New(v), nil
}
