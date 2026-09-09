package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (c *Client) MustCreateTag(
	project int64,
	name string,
	reference string,
	message string,
) *tag.Tag {
	result, e := c.CreateTag(project, name, reference, message)
	errors.PanicOnError(e)

	return result
}
