package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) UpdateIssue(
	organization string,
	identifier string,
	status string,
	assignedTo string,
) (*response.Issue, error) {
	type body struct {
		Status     string `json:"status,omitempty"`
		AssignedTo string `json:"assignedTo,omitempty"`
	}
	var result response.Issue

	if e := c.basic.Put(
		fmt.Sprintf("organizations/%s/issues/%s", organization, identifier),
		body{Status: status, AssignedTo: assignedTo},
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
