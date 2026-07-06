package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustSearchIssues(
	organization string,
	query string,
	project string,
	limit int,
	cursor string,
) []response.Issue {
	result, e := c.SearchIssues(organization, query, project, limit, cursor)
	errors.PanicOnError(e)

	return result
}
