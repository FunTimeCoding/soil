package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCommitActions(
	project int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) *commit.Commit {
	result, e := c.CommitActions(project, branch, message, v)
	errors.PanicOnError(e)

	return result
}
