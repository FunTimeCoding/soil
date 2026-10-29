package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedTenant(name string) {
	v := netbox.NewTenantWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	c.tenants = append(c.tenants, tenant.New(v))
}
