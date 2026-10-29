package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ModuleBays() ([]*module_bay.Bay, error) {
	result, _, e := c.client.DcimAPI.DcimModuleBaysList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return module_bay.NewSlice(result.Results), nil
}
