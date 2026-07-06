package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCommit(
	project int64,
	branch string,
	text string,
	path string,
	content string,
	update bool,
) *gitlab.Commit {
	result, e := c.Commit(project, branch, text, path, content, update)
	errors.PanicOnError(e)

	return result
}
