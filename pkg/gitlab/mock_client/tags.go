package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (c *Client) Tags(
	_ context.Context,
	_ int64,
) ([]*tag.Tag, error) {
	return c.tags, nil
}
