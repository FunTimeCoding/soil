package hetzner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
)

func (c *Client) Zones() []*zone.Zone {
	result, e := c.client.Zone.All(c.context)
	errors.PanicOnError(e)

	return zone.NewSlice(result)
}
