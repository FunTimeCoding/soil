package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) IssueLink(key string) string {
	return issue.BuildLink(c.locator, key)
}
