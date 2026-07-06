package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) SearchEvents(
	organization string,
	query string,
	project string,
	limit int,
	cursor string,
) ([]response.Event, error) {
	q := map[string]string{}

	if query != "" {
		q["query"] = query
	}

	if project != "" {
		q["project"] = project
	}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	if cursor != "" {
		q["cursor"] = cursor
	}

	b, e := c.basic.Get(
		fmt.Sprintf("organizations/%s/events", organization),
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
