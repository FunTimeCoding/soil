package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) MustTenant(n string) *tenant.Tenant {
	result, e := c.Tenant(n)
	errors.PanicOnError(e)

	return result
}
