package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustFileExists(
	x context.Context,
	p *project.Project,
	branch string,
	file string,
) bool {
	result, e := c.FileExists(x, p, branch, file)
	errors.PanicOnError(e)

	return result
}
