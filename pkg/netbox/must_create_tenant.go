package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) MustCreateTenant(name string) *tenant.Tenant {
	result, e := c.CreateTenant(name)
	errors.PanicOnError(e)

	return result
}
