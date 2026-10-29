package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/issue"

func (c *Client) wrappedIssues() ([]*issue.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []*issue.Issue

	for _, i := range c.issues {
		result = append(result, issue.New(i))
	}

	return result, nil
}
