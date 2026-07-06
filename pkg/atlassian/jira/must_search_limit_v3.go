package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSearchLimitV3(
	limit int,
	query string,
	a ...any,
) []*response.Issue {
	result, e := c.SearchLimitV3(limit, query, a...)
	errors.PanicOnError(e)

	return result
}
