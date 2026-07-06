package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) MustTeams(organization string) []response.Team {
	result, e := c.Teams(organization)
	errors.PanicOnError(e)

	return result
}
