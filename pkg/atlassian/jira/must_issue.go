package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustIssue(key string) *issue.Issue {
	result, e := c.Issue(key)
	errors.PanicOnError(e)

	return result
}
