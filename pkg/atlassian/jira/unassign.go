package jira

func (c *Client) Unassign(key string) error {
	r, e := c.client.Issue.UpdateAssigneeWithContext(c.context, key, nil)

	return WrapError(r, e)
}
