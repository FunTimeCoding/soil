package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/module_type_profile"

func (c *Client) ModuleTypeProfiles() ([]*module_type_profile.Profile, error) {
	return nil, c.Failure
}
