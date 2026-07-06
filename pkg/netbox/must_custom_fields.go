package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
)

func (c *Client) MustCustomFields() []*custom_field.Field {
	result, e := c.CustomFields()
	errors.PanicOnError(e)

	return result
}
