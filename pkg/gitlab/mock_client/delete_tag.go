package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"slices"
)

func (c *Client) DeleteTag(
	_ context.Context,
	_ int64,
	name string,
) error {
	if c.failure != nil {
		return c.failure
	}

	c.tags = slices.DeleteFunc(
		c.tags,
		func(t *tag.Tag) bool { return t.Name == name },
	)

	return nil
}
