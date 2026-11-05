package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"slices"
)

func (c *Client) DeleteConfigurationContext(name string) error {
	if c.Failure != nil {
		return c.Failure
	}

	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return e
	}

	c.contexts = slices.DeleteFunc(
		c.contexts,
		func(x *configuration_context.Context) bool { return x == v },
	)
	c.record("DeleteConfigurationContext", name)

	return nil
}
