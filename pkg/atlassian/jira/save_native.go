package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) SaveNative(i *jira.Issue) (*issue.Issue, error) {
	o, e := c.IssueOption()

	if e != nil {
		return nil, e
	}

	result, _, f := c.client.Issue.UpdateWithContext(c.context, i)

	if f != nil {
		return nil, f
	}

	return c.enrichOne(issue.New(result, o)), nil
}
