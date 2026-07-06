package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/module_type_profile"
)

func (c *Client) MustModuleTypeProfiles() []*module_type_profile.Profile {
	result, e := c.ModuleTypeProfiles()
	errors.PanicOnError(e)

	return result
}
