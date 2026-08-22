package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) AddLink(link *jira.IssueLink) error {
	_, e := c.client.Issue.AddLinkWithContext(c.context, link)

	return wrapError(e)
}
