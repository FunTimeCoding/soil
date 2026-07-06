package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/system_number"
)

func (c *Client) MustSystemNumbers() []*system_number.Number {
	result, e := c.SystemNumbers()
	errors.PanicOnError(e)

	return result
}
