package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/configuration_context"

func (c *Client) SetConfigurationContextPayload(
	name string,
	payload map[string]any,
) (*configuration_context.Context, error) {
	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return nil, e
	}

	v.Payload = payload
	c.record("SetConfigurationContextPayload", name)

	return v, nil
}
