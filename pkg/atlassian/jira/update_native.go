package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) UpdateNative(i *jira.Issue) error {
	_, r, e := c.client.Issue.UpdateWithContext(c.context, i)

	return WrapError(r, e)
}
