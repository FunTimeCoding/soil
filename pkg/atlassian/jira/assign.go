package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Assign(
	key string,
	u *jira.User,
) error {
	r, e := c.client.Issue.UpdateAssigneeWithContext(c.context, key, u)

	return WrapError(r, e)
}
