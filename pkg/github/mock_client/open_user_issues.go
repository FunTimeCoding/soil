package mock_client

import "github.com/funtimecoding/soil/pkg/github/issue"

func (c *Client) OpenUserIssues() ([]*issue.Issue, error) {
	return nil, c.Failure
}
