package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) GetTags(context.Context) ([]*client.Tag, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.TagList, nil
}
