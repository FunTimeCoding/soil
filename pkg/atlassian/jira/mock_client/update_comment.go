package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) UpdateComment(
	_ string,
	_ *jira.Comment,
) error {
	return nil
}
