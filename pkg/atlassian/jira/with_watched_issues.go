package jira

func WithWatchedIssues() Option {
	return func(c *Client) {
		c.watchedIssues = true
	}
}
