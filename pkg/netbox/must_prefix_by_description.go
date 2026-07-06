package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
)

func (c *Client) MustPrefixByDescription(d string) *prefix.Prefix {
	result, e := c.PrefixByDescription(d)
	errors.PanicOnError(e)

	return result
}
