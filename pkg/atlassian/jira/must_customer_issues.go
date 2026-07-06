package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/customer"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustCustomerIssues(all bool) []*customer.Issue {
	result, e := c.CustomerIssues(all)
	errors.PanicOnError(e)

	return result
}
