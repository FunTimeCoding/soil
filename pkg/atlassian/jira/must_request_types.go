package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustRequestTypes(
	desk int,
	group int,
) *models.ProjectRequestTypePageScheme {
	result, e := c.RequestTypes(desk, group)
	errors.PanicOnError(e)

	return result
}
