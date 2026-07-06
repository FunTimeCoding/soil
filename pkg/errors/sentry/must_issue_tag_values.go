package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustIssueTagValues(
	organization string,
	identifier string,
	tag string,
	limit int,
) []response.TagValue {
	result, e := c.IssueTagValues(organization, identifier, tag, limit)
	errors.PanicOnError(e)

	return result
}
