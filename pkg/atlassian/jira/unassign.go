package jira

func (c *Client) Unassign(key string) error {
	_, e := c.client.Issue.UpdateAssigneeWithContext(c.context, key, nil)

	return e
}
