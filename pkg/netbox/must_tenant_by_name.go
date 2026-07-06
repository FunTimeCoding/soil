package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) MustTenantByName(n string) *tenant.Tenant {
	result, e := c.TenantByName(n)
	errors.PanicOnError(e)

	return result
}
