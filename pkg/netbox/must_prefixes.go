package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
)

func (c *Client) MustPrefixes() []*prefix.Prefix {
	result, e := c.Prefixes()
	errors.PanicOnError(e)

	return result
}
