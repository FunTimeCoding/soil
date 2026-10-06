package sentry

import (
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

	var result []response.Event

	if e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s/events",
			organization,
			identifier,
		),
		q,
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
