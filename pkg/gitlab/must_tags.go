package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (c *Client) MustTags(project int64) []*tag.Tag {
	result, e := c.Tags(project)
	errors.PanicOnError(e)

	return result
}
