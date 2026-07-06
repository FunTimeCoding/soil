package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustWatchedIssues() []*issue.Issue {
	result, e := c.WatchedIssues()
	errors.PanicOnError(e)

	return result
}
