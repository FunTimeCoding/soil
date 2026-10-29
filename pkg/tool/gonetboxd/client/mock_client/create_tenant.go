package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateTenant(_ context.Context, _ string) (client.Tenant, error) {
	return client.Tenant{}, c.Failure
}
