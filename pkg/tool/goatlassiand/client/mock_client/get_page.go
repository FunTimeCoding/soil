package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetPage(
	_ context.Context,
	_ string,
) (*client.ConfluencePageDetail, error) {
	return nil, c.Failure
}
