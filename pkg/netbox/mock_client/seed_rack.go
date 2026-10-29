package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedRack(
	name string,
	place string,
	height int32,
) {
	v := netbox.NewRackWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	v.SetSite(*netbox.NewBriefSite(0, "", place, place, place))
	v.SetUHeight(height)
	c.racks = append(c.racks, rack.New(v))
}
