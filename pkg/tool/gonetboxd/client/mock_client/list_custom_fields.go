package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListCustomFields(_ context.Context) ([]*client.CustomField, error) {
	return nil, c.Failure
}
