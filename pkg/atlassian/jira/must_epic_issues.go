package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustEpicIssues(name string) []*issue.Issue {
	result, e := c.EpicIssues(name)
	errors.PanicOnError(e)

	return result
}
