package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/issue"

func (c *Client) Issues(
	_ string,
	_ string,
	_ string,
) ([]*issue.Issue, error) {
	return c.wrappedIssues()
}
