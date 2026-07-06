package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
)

func (c *Client) RackTypes() ([]*rack_type.Type, error) {
	result, _, e := c.client.DcimAPI.DcimRackTypesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return rack_type.NewSlice(result.Results), nil
}
