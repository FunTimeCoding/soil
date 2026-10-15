package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
)

func (c *Client) MustCustomFieldByName(n string) *custom_field.Field {
	result, e := c.CustomFieldByName(n)
	errors.PanicOnError(e)

	return result
}
