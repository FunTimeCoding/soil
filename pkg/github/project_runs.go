package github

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/google/go-github/v88/github"
)

func (c *Client) ProjectRuns(
	owner string,
	name string,
) ([]*run.Run, error) {
	var result []*run.Run
	o := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{
			PerPage: 100, // Cannot go higher than 100
		},
	}

	for {
		page, r, e := c.client.Actions.ListRepositoryWorkflowRuns(
			c.context,
			owner,
			name,
			o,
		)

		if e != nil {
			return nil, e
		}

		result = append(result, run.NewSlice(page.WorkflowRuns)...)

		if r.NextPage == 0 {
			break
		}

		o.Page = r.NextPage
	}

	for _, s := range result {
		s.Validate()
	}

	return result, nil
}
