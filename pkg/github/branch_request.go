package github

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/google/go-github/v90/github"
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
		return nil, not_found.Format(
			"repository not found: %s/%s",
			owner,
			repository,
		)
	}

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, not_found.Format(
			"no pull request for branch %s in %s/%s",
			branch,
			owner,
			repository,
		)
	}

	return pull_request.New(result[0]), nil
}
