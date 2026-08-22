package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) RackByName(n string) (*rack.Rack, error) {
	result, e := c.RacksByName(n)

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, not_found.New("rack", n)
	}

	if len(result) > 1 {
		return nil, ambiguous.Format(
			"expected 1 rack named %s, got %d",
			n,
			len(result),
		)
	}

	return result[0], nil
}
