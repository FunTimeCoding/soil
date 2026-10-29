package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) GetConfigurationContext(_ context.Context, _ string) (client.ConfigurationContext, error) {
	return client.ConfigurationContext{}, c.Failure
}
