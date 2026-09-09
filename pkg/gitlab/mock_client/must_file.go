package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
)

func (c *Client) MustFile(
	project int64,
	branch string,
	name string,
) *file.File {
	result, e := c.File(project, branch, name)
	errors.PanicOnError(e)

	return result
}
