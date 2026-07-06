package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
)

func (c *Client) ModuleTypes() ([]*module_type.Type, error) {
	result, _, e := c.client.DcimAPI.DcimModuleTypesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return module_type.NewSlice(result.Results), nil
}
