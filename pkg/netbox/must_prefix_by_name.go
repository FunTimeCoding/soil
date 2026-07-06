package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
)

func (c *Client) MustPrefixByName(n string) *prefix.Prefix {
	result, e := c.PrefixByName(n)
	errors.PanicOnError(e)

	return result
}
