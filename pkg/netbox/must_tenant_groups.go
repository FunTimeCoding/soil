package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
)

func (c *Client) MustTenantGroups() []*tenant_group.Group {
	result, e := c.TenantGroups()
	errors.PanicOnError(e)

	return result
}
