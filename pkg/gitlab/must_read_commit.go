package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
)

func (c *Client) MustReadCommit(
	project int64,
	sha string,
) *commit.Commit {
	result, e := c.ReadCommit(project, sha)
	errors.PanicOnError(e)

	return result
}
