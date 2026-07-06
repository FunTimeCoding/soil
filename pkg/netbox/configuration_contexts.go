package netbox

import "github.com/funtimecoding/soil/pkg/netbox/configuration_context"

func (c *Client) ConfigurationContexts() ([]*configuration_context.Context, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasConfigContextsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return configuration_context.NewSlice(result.Results), nil
}
