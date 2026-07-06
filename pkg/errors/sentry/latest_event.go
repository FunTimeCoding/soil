package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) LatestEvent(
	organization string,
	issueIdentifier string,
) (*response.Event, error) {
	b, e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s/events/latest",
			organization,
			issueIdentifier,
		),
		nil,
	)

	if e != nil {
		return nil, e
	}

	var result response.Event

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
