package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tenant"

func (c *Client) CreateTenant(_ string) (*tenant.Tenant, error) {
	return nil, nil
}
