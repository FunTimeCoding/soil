package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue/issue_enricher"

func (c *Client) Set(e *issue_enricher.Enricher) *Client {
	c.enricher = e

	return c
}
