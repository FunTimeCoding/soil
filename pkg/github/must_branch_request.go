package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
)

func (c *Client) MustBranchRequest(
	owner string,
	repository string,
	branch string,
) *pull_request.Request {
	result, e := c.BranchRequest(owner, repository, branch)
	errors.PanicOnError(e)

	return result
}
