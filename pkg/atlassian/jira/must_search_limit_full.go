package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSearchLimitFull(
	limit int,
	query string,
	a ...any,
) []*issue.Issue {
	result, e := c.SearchLimitFull(limit, query, a...)
	errors.PanicOnError(e)

	return result
}
