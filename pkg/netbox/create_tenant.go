package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTenant(name string) (*tenant.Tenant, error) {
	q := netbox.NewTenantRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.TenancyAPI.TenancyTenantsCreate(
		c.context,
	).TenantRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Tenants = nil

	return tenant.New(result), nil
}
