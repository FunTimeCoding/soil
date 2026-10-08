package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/update_request"
)

func (c *Client) UpdateIssue(
	organization string,
	identifier string,
	status string,
	assignedTo string,
) (*response.Issue, error) {
	var result response.Issue

	if e := c.basic.Put(
		fmt.Sprintf("organizations/%s/issues/%s", organization, identifier),
		update_request.New(status, assignedTo),
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
