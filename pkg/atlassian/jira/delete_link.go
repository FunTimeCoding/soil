package jira

func (c *Client) DeleteLink(identifier string) error {
	r, e := c.client.Issue.DeleteLinkWithContext(c.context, identifier)

	return WrapError(r, e)
}
