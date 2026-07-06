package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustCreateCustomerIssue(
	desk int,
	requestType int,
	summary string,
	description string,
) *models.CustomerRequestScheme {
	result, e := c.CreateCustomerIssue(
		desk,
		requestType,
		summary,
		description,
	)
	errors.PanicOnError(e)

	return result
}
