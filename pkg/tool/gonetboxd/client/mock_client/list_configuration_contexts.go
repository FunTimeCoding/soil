package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListConfigurationContexts(_ context.Context) ([]*client.ConfigurationContext, error) {
	return nil, c.Failure
}
