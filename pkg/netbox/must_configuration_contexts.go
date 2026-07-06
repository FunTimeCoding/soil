package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
)

func (c *Client) MustConfigurationContexts() []*configuration_context.Context {
	result, e := c.ConfigurationContexts()
	errors.PanicOnError(e)

	return result
}
