package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Event(
	organization string,
	project string,
	identifier string,
) (*response.Event, error) {
	var result response.Event

	if e := c.basic.Get(
		fmt.Sprintf(
			"projects/%s/%s/events/%s",
			organization,
			project,
			identifier,
		),
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
