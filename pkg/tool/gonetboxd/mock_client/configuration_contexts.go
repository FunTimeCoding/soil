package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/configuration_context"

func (c *Client) ConfigurationContexts() ([]*configuration_context.Context, error) {
	return c.contexts, nil
}
