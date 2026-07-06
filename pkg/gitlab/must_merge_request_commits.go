package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
)

func (c *Client) MustMergeRequestCommits(
	project int64,
	request int64,
) []*commit.Commit {
	result, e := c.MergeRequestCommits(project, request)
	errors.PanicOnError(e)

	return result
}
