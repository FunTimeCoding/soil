package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateConfigurationContext(_ context.Context, _ string, _ client.UpdateConfigurationContextRequest) (client.ConfigurationContext, error) {
	return client.ConfigurationContext{}, c.Failure
}
