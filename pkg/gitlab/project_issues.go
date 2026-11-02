package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ProjectIssues(
	x context.Context,
	project int64,
) ([]*issue.Issue, error) {
	result, _, e := c.client.Issues.ListProjectIssues(
		project,
		&gitlab.ListProjectIssuesOptions{
			ListOptions: constant.DefaultListOptions,
		},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return issue.NewSlice(result), nil
}
