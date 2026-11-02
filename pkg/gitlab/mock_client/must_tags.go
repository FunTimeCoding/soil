package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (c *Client) MustTags(
	_ context.Context,
	_ int64,
) []*tag.Tag {
	return c.tags
}
