package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustIssueEvents(
	organization string,
	identifier string,
	query string,
	limit int,
	cursor string,
) []response.Event {
	result, e := c.IssueEvents(organization, identifier, query, limit, cursor)
	errors.PanicOnError(e)

	return result
}
