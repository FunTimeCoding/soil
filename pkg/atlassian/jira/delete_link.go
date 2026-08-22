package jira

func (c *Client) DeleteLink(identifier string) error {
	_, e := c.client.Issue.DeleteLinkWithContext(c.context, identifier)

	return wrapError(e)
}
