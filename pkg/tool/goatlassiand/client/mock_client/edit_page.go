package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) EditPage(
	_ context.Context,
	_ string,
	_ string,
	_ string,
	_ string,
	_ string,
	_ bool,
) (*client.ConfluencePageDetail, error) {
	return nil, c.Failure
}
