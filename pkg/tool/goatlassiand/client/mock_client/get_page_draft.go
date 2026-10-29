package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetPageDraft(
	_ context.Context,
	_ string,
) (*client.ConfluencePageDetail, error) {
	return nil, c.Failure
}
