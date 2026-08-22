package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) UpdateNative(i *jira.Issue) error {
	_, _, e := c.client.Issue.UpdateWithContext(c.context, i)

	return wrapError(e)
}
