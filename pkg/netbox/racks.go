package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) Racks() ([]*rack.Rack, error) {
	if len(c.cache.Racks) != 0 {
		return c.cache.Racks, nil
	}

	result, _, e := c.client.DcimAPI.DcimRacksList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Racks = rack.NewSlice(result.Results)

	return c.cache.Racks, nil
}
