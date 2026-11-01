package mock_client

import "github.com/funtimecoding/soil/pkg/github/issue"

func (c *Client) SearchIssue(_ string, _ ...any) ([]*issue.Issue, error) {
	return nil, c.Failure
}
