package netbox

import "github.com/funtimecoding/soil/pkg/netbox/module_type_profile"

func (c *Client) ModuleTypeProfiles() ([]*module_type_profile.Profile, error) {
	result, _, e := c.client.DcimAPI.DcimModuleTypeProfilesList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return module_type_profile.NewSlice(result.Results), nil
}
