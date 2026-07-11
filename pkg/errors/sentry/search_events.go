package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"net/url"
	"strconv"
)

func (c *Client) SearchEvents(
	organization string,
	query string,
	project string,
	limit int,
	cursor string,
) ([]response.EventRow, error) {
	q := url.Values{}

	for _, f := range constant.EventFields {
		q.Add("field", f)
	}

	if query != "" {
		q.Set("query", query)
	}

	if project != "" {
		q.Set("project", project)
	}

	if limit > 0 {
		q.Set("per_page", strconv.Itoa(limit))
	}

	if cursor != "" {
		q.Set("cursor", cursor)
	}

	b, e := c.basic.GetValues(
		fmt.Sprintf("organizations/%s/events", organization),
		q,
	)

	if e != nil {
		return nil, e
	}

	var result response.EventSearch

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result.Rows, nil
}
