package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) Assign(
	_ string,
	_ *jira.User,
) error {
	return nil
}
