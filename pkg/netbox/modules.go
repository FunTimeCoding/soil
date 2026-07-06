package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/module"
)

func (c *Client) Modules() ([]*module.Module, error) {
	result, _, e := c.client.DcimAPI.DcimModulesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return module.NewSlice(result.Results), nil
}
