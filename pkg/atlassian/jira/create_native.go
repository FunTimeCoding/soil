package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) CreateNative(i *jira.Issue) (*issue.Issue, error) {
	result, _, e := c.client.Issue.CreateWithContext(c.context, i)

	if e != nil {
		return nil, e
	}

	v, f := c.Issue(result.Key)

	if f != nil {
		return nil, f
	}

	return c.enrichOne(v), nil
}
