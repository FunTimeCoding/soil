package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateCustomField(_ context.Context, _ string, _ string, _ string) (client.CustomField, error) {
	return client.CustomField{}, c.Failure
}
