package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateManufacturer(name string) (*manufacturer.Manufacturer, error) {
	q := netbox.NewManufacturerRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.DcimAPI.DcimManufacturersCreate(
		c.context,
	).ManufacturerRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Manufacturers = nil

	return manufacturer.New(result), nil
}
