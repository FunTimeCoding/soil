package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) AddLink(_ *jira.IssueLink) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
