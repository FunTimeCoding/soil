package sentry

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteIssue(
	organization string,
	identifier string,
) {
	errors.PanicOnError(c.DeleteIssue(organization, identifier))
}
