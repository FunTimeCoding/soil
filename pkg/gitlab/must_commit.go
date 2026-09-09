package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
)

func (c *Client) MustCommit(
	project int64,
	branch string,
	text string,
	path string,
	content string,
	update bool,
) *commit.Commit {
	result, e := c.Commit(project, branch, text, path, content, update)
	errors.PanicOnError(e)

	return result
}
