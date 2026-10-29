package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) SearchPages(
	_ context.Context,
	_ string,
) ([]*client.ConfluencePage, error) {
	return nil, c.Failure
}
