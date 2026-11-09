package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) MustRack(n string) *rack.Rack {
	result, e := c.Rack(n)
	errors.PanicOnError(e)

	return result
}
