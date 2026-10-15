package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_request"
)

func (c *Client) SetConfigurationContextAssignment(
	name string,
	a *configuration_assignment.Assignment,
) (*configuration_context.Context, error) {
	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return nil, e
	}

	q := configuration_request.New()

	if f := c.assignmentFields(q, a); f != nil {
		return nil, f
	}

	return c.updateConfigurationContext(v, q)
}
