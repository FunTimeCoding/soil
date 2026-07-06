package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustDesks() *models.ServiceDeskPageScheme {
	result, e := c.Desks()
	errors.PanicOnError(e)

	return result
}
