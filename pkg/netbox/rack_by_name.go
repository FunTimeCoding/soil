package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) RackByName(n string) (*rack.Rack, error) {
	result, e := c.RacksByName(n)

	if e != nil {
		return nil, e
	}

	if len(result) != 1 {
		return nil, fmt.Errorf(
			"expected 1 rack named %s, got %d",
			n,
			len(result),
		)
	}

	return result[0], nil
}
