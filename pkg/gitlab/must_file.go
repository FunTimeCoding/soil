package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustFile(
	project int64,
	branch string,
	name string,
) *gitlab.File {
	result, e := c.File(project, branch, name)
	errors.PanicOnError(e)

	return result
}
