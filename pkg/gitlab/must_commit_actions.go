package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCommitActions(
	project int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) *gitlab.Commit {
	result, e := c.CommitActions(project, branch, message, v)
	errors.PanicOnError(e)

	return result
}
