package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustCustomerIssuesBasic() {
	errors.PanicOnError(c.CustomerIssuesBasic())
}
