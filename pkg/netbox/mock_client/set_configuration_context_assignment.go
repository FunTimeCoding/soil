package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
)

func (c *Client) SetConfigurationContextAssignment(
	name string,
	a *configuration_assignment.Assignment,
) (*configuration_context.Context, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return nil, e
	}

	assign(v, a)
	c.record("SetConfigurationContextAssignment", name)

	return v, nil
}
