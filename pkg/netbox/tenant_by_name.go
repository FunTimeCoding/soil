package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) TenantByName(n string) (*tenant.Tenant, error) {
	result, e := c.Tenants()

	if e != nil {
		return nil, e
	}

	for _, t := range result {
		if t.Name == n {
			return t, nil
		}
	}

	return nil, not_found.New("tenant", n)
}
