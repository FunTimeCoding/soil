package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/issue"

func (c *Client) TrackedIssues() ([]*issue.Issue, error) {
	return c.wrappedIssues()
}
