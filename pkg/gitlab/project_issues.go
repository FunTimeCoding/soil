package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ProjectIssues(project int64) ([]*issue.Issue, error) {
	result, _, e := c.client.Issues.ListProjectIssues(
		project,
		&gitlab.ListProjectIssuesOptions{
			ListOptions: constant.DefaultListOptions,
		},
	)

	if e != nil {
		return nil, e
	}

	return issue.NewSlice(result), nil
}
