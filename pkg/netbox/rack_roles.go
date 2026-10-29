package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) RackRoles() ([]*rack_role.Role, error) {
	result, _, e := c.client.DcimAPI.DcimRackRolesList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return rack_role.NewSlice(result.Results), nil
}
