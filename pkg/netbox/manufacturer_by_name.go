package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
)

func (c *Client) ManufacturerByName(n string) (*manufacturer.Manufacturer, error) {
	result, e := c.Manufacturers()

	if e != nil {
		return nil, e
	}

	for _, m := range result {
		if m.Name == n {
			return m, nil
		}
	}

	return nil, fmt.Errorf("manufacturer not found: %s", n)
}
