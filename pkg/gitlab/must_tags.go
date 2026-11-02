package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (c *Client) MustTags(
	x context.Context,
	project int64,
) []*tag.Tag {
	result, e := c.Tags(x, project)
	errors.PanicOnError(e)

	return result
}
