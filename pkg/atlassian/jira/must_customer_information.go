package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustCustomerInformation() *models.InfoScheme {
	result, e := c.CustomerInformation()
	errors.PanicOnError(e)

	return result
}
