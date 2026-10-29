package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tenant_group"

func (c *Client) TenantGroups() ([]*tenant_group.Group, error) {
	return nil, c.Failure
}
