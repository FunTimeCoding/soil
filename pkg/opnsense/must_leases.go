package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/lease"
)

func (c *Client) MustLeases(phrase string) []*lease.Lease {
	result, e := c.Leases(phrase)
	errors.PanicOnError(e)

	return result
}
