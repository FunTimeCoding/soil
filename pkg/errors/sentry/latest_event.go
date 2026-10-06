package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) LatestEvent(
	organization string,
	issueIdentifier string,
) (*response.Event, error) {
	var result response.Event

	if e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s/events/latest",
			organization,
			issueIdentifier,
		),
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
