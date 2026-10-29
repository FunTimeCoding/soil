package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) TenantByName(n string) (*tenant.Tenant, error) {
	for _, t := range c.tenants {
		if t.Name == n {
			return t, nil
		}
	}

	return nil, not_found.New(constant.TenantKind, n)
}
