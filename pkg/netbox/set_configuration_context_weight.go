package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_request"
)

func (c *Client) SetConfigurationContextWeight(
	name string,
	weight int32,
) (*configuration_context.Context, error) {
	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return nil, e
	}

	q := configuration_request.New()
	q.Weight = &weight

	return c.updateConfigurationContext(v, q)
}
