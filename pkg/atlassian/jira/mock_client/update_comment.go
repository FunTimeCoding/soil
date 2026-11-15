package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) UpdateComment(
	_ string,
	_ *jira.Comment,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
