package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedRawRack(v *netbox.Rack) {
	v.SetId(c.nextIdentifier())
	c.racks = append(c.racks, rack.New(v))
}
