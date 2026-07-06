package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) WatchedIssues() ([]*issue.Issue, error) {
	return c.SearchFull("issue in watchedIssues() ORDER BY key ASC")
}
