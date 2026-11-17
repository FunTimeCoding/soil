package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) RenameDeviceRole(
	name string,
	newName string,
) (*device_role.Role, error) {
	r, e := c.DeviceRole(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableDeviceRoleRequest()
	q.SetName(newName)
	q.SetSlug(slug(newName))
	result, _, f := c.client.DcimAPI.DcimDeviceRolesPartialUpdate(
		c.context,
		r.Identifier,
	).PatchedWritableDeviceRoleRequest(*q).Execute()

	if f != nil {
		return nil, refusal.Classify(f)
	}

	c.cache.DeviceRoles = nil

	return device_role.New(result), nil
}
