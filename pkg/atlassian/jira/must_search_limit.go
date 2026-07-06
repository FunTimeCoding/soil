package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSearchLimit(
	limit int,
	query string,
	a ...any,
) []*issue.Issue {
	result, e := c.SearchLimit(limit, query, a...)
	errors.PanicOnError(e)

	return result
}
