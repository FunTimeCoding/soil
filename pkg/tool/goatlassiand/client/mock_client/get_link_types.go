package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetLinkTypes(_ context.Context) ([]client.LinkType, error) {
	return nil, c.Failure
}
