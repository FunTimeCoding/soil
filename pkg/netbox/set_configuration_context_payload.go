package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_request"
)

func (c *Client) SetConfigurationContextPayload(
	name string,
	payload map[string]any,
) (*configuration_context.Context, error) {
	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return nil, e
	}

	q := configuration_request.New()
	q.Payload = &payload

	return c.updateConfigurationContext(v, q)
}
