package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) Manufacturers() ([]*manufacturer.Manufacturer, error) {
	if len(c.cache.Manufacturers) != 0 {
		return c.cache.Manufacturers, nil
	}

	result, _, e := c.client.DcimAPI.DcimManufacturersList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	c.cache.Manufacturers = manufacturer.NewSlice(result.Results)

	return c.cache.Manufacturers, nil
}
