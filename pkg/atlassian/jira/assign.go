package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Assign(
	key string,
	u *jira.User,
) error {
	_, e := c.client.Issue.UpdateAssigneeWithContext(c.context, key, u)

	return wrapError(e)
}
