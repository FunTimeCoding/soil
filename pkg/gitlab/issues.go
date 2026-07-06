package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Issues() ([]*issue.Issue, error) {
	result, _, e := c.client.Issues.ListIssues(
		&gitlab.ListIssuesOptions{
			ListOptions: constant.DefaultListOptions,
		},
	)

	if e != nil {
		return nil, e
	}

	return issue.NewSlice(result), nil
}
