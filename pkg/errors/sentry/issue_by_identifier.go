package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) IssueByIdentifier(
	organization string,
	identifier string,
) (*response.Issue, error) {
	b, e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s",
			organization,
			identifier,
		),
		nil,
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
