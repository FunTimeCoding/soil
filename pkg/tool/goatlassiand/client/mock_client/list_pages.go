package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) ListPages(
	_ context.Context,
	_ string,
	_ string,
) ([]*client.ConfluencePage, error) {
	return nil, c.Failure
}
