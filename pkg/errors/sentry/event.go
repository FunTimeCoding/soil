package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Event(
	organization string,
	project string,
	identifier string,
) (*response.Event, error) {
	b, e := c.basic.Get(
		fmt.Sprintf(
			"projects/%s/%s/events/%s",
			organization,
			project,
			identifier,
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
