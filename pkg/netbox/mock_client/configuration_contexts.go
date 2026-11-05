package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/configuration_context"

func (c *Client) ConfigurationContexts() ([]*configuration_context.Context, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.contexts, nil
}
