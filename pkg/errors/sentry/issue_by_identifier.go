package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) IssueByIdentifier(
	organization string,
	identifier string,
) (*response.Issue, error) {
	var result response.Issue

	if e := c.basic.Get(
		fmt.Sprintf("organizations/%s/issues/%s", organization, identifier),
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
