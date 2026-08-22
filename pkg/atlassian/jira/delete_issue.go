package jira

func (c *Client) DeleteIssue(key string) error {
	_, e := c.client.Issue.Delete(key)

	return wrapError(e)
}
