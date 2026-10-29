package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) GetTag(
	_ context.Context,
	_ *client.GetTagParams,
) (*client.SourceTypeTag, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.SourceTypeTag{}, nil
}
