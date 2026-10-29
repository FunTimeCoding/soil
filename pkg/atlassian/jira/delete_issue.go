package jira

func (c *Client) DeleteIssue(key string) error {
	r, e := c.client.Issue.Delete(key)

	return WrapError(r, e)
}
