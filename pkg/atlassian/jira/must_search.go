package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSearch(
	query string,
	a ...any,
) []*issue.Issue {
	result, e := c.Search(query, a...)
	errors.PanicOnError(e)

	return result
}
