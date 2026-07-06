package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) Issue(key string) (*issue.Issue, error) {
	o, e := c.IssueOption()

	if e != nil {
		return nil, e
	}

	result, _, f := c.client.Issue.Get(key, &jira.GetQueryOptions{})

	if f != nil {
		return nil, f
	}

	return issue.New(result, o), nil
}
