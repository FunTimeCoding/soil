package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListTenants(_ context.Context) ([]*client.Tenant, error) {
	return nil, c.Failure
}
