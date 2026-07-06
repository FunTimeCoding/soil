package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/customer"
	"github.com/funtimecoding/soil/pkg/integers"
)

func (c *Client) CreateCustomerIssue(
	desk int,
	requestType int,
	summary string,
	description string,
) (*models.CustomerRequestScheme, error) {
	result, _, e := c.service.Request.Create(
		c.context,
		&models.CreateCustomerRequestPayloadScheme{
			ServiceDeskID: integers.ToString(desk),
			RequestTypeID: integers.ToString(requestType),
			RequestFieldValues: map[string]any{
				customer.SummaryField:     summary,
				customer.DescriptionField: description,
			},
		},
	)

	return result, e
}
