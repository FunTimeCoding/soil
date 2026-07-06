package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) enrichMany(v []*issue.Issue) []*issue.Issue {
	if c.enricher == nil {
		return v
	}

	return c.enricher.Run(v)
}
