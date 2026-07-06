package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeletePackage(
	project int64,
	repository int64,
) {
	errors.PanicOnError(c.DeletePackage(project, repository))
}
