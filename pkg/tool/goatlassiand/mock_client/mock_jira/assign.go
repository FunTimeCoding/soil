package mock_jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Assign(
	_ string,
	_ *jira.User,
) error {
	return nil
}
