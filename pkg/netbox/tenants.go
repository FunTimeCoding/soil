package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) Tenants() ([]*tenant.Tenant, error) {
	if len(c.cache.Tenants) != 0 {
		return c.cache.Tenants, nil
	}

	result, _, e := c.client.TenancyAPI.TenancyTenantsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	c.cache.Tenants = tenant.NewSlice(result.Results)

	return c.cache.Tenants, nil
}
