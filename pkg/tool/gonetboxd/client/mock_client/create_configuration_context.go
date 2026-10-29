package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateConfigurationContext(_ context.Context, _ client.CreateConfigurationContextRequest) (client.ConfigurationContext, error) {
	return client.ConfigurationContext{}, c.Failure
}
