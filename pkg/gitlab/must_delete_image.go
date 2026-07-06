package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteImage(
	project int64,
	repository int64,
	tag string,
) {
	errors.PanicOnError(c.DeleteImage(project, repository, tag))
}
