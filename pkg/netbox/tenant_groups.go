package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
)

func (c *Client) TenantGroups() ([]*tenant_group.Group, error) {
	result, _, e := c.client.TenancyAPI.TenancyTenantGroupsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return tenant_group.NewSlice(result.Results), nil
}
