package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/issue"

func (c *Client) ProjectIssues(_ int64) ([]*issue.Issue, error) {
	return nil, c.failure
}
