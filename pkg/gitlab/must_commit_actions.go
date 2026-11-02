package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustCommitActions(
	x context.Context,
	project int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) *commit.Commit {
	result, e := c.CommitActions(x, project, branch, message, v)
	errors.PanicOnError(e)

	return result
}
