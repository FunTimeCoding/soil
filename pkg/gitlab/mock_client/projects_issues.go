package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/issue"

func (c *Client) ProjectsIssues() ([]*issue.Issue, error) {
	return nil, c.failure
}
