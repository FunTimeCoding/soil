package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) enrichOne(i *issue.Issue) *issue.Issue {
	if c.enricher == nil {
		return i
	}

	return c.enricher.Run([]*issue.Issue{i})[0]
}
