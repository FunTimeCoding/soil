package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteTag(
	project int64,
	name string,
) {
	errors.PanicOnError(c.DeleteTag(project, name))
}
