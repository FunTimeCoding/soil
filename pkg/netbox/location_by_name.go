package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/location"
)

func (c *Client) LocationByName(n string) (*location.Location, error) {
	result, e := c.Locations()

	if e != nil {
		return nil, e
	}

	for _, l := range result {
		if l.Name == n {
			return l, nil
		}
	}

	return nil, not_found.New("location", n)
}
