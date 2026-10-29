package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) AddLink(link *jira.IssueLink) error {
	r, e := c.client.Issue.AddLinkWithContext(c.context, link)

	return WrapError(r, e)
}
