package jira

func (c *Client) WatchedIssueKeys() ([]string, error) {
	return c.SearchKeys("issue IN watchedIssues() ORDER BY key ASC")
}
