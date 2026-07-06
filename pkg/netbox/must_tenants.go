package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) MustTenants() []*tenant.Tenant {
	result, e := c.Tenants()
	errors.PanicOnError(e)

	return result
}
