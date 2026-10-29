package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) ListSpaces(_ context.Context) ([]*client.ConfluenceSpace, error) {
	return nil, c.Failure
}
