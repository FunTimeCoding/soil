package github

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/google/go-github/v89/github"
)

func (c *Client) BranchRequest(
	owner string,
	repository string,
	branch string,
) (*pull_request.Request, error) {
	result, r, e := c.client.PullRequests.List(
		c.context,
		owner,
		repository,
		&github.PullRequestListOptions{
			State:       constant.All,
			Head:        join.Colon(owner, branch),
			ListOptions: github.ListOptions{PerPage: 1},
		},
	)

	if r != nil && r.StatusCode == 404 {
		return nil, fmt.Errorf(
			"repository not found: %s/%s: %w",
			owner,
			repository,
			constant.ErrorNotFound,
		)
	}

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"no pull request for branch %s in %s/%s: %w",
			branch,
			owner,
			repository,
			constant.ErrorNotFound,
		)
	}

	return pull_request.New(result[0]), nil
}
