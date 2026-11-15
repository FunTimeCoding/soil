package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) Assign(
	_ string,
	_ *jira.User,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
