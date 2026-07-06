package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/location"
)

func (c *Client) Locations() ([]*location.Location, error) {
	result, _, e := c.client.DcimAPI.DcimLocationsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return location.NewSlice(result.Results), nil
}
