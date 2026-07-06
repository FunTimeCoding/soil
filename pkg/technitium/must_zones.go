package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/zone"
)

func (c *Client) MustZones() []*zone.Zone {
	result, e := c.Zones()
	errors.PanicOnError(e)

	return result
}
