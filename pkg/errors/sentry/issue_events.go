package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) IssueEvents(
	organization string,
	identifier string,
	query string,
	limit int,
	cursor string,
) ([]response.Event, error) {
	q := map[string]string{"full": "1"}

	if query != "" {
		q["query"] = query
	}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	if cursor != "" {
		q["cursor"] = cursor
	}

	b, e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s/events",
			organization,
			identifier,
		),
		q,
	)

	if e != nil {
		return nil, e
	}

	var result []response.Event

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
