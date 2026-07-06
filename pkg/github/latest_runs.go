package github

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/google/go-github/v88/github"
)

func (c *Client) LatestRuns(owner, repo string) ([]*run.Run, error) {
	page, _, e := c.client.Actions.ListRepositoryWorkflowRuns(
		c.context,
		owner,
		repo,
		&github.ListWorkflowRunsOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return nil, e
	}

	return run.NewSlice(page.WorkflowRuns), nil
}
