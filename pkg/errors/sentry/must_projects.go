package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustProjects() []response.Project {
	result, e := c.Projects()
	errors.PanicOnError(e)

	return result
}
