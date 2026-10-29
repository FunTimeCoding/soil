package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) User() (*jira.User, error) {
	result, r, e := c.client.User.GetSelf()

	return result, WrapError(r, e)
}
