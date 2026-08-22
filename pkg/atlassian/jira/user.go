package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) User() (*jira.User, error) {
	result, _, e := c.client.User.GetSelf()

	return result, wrapError(e)
}
