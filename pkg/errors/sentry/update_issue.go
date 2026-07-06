package sentry

import (
	"encoding/json"
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
	b, e := c.basic.Put(
		fmt.Sprintf(
			"organizations/%s/issues/%s",
			organization,
			identifier,
		),
		body{Status: status, AssignedTo: assignedTo},
	)

	if e != nil {
		return nil, e
	}

	var result response.Issue

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
