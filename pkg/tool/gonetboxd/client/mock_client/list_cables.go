package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListCables(_ context.Context) ([]*client.Cable, error) {
	return nil, c.Failure
}
