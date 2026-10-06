package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
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

	q.Set("sort", constant.SortNewestFirst)

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

	var result response.EventSearch

	if e := c.basic.GetValues(
		fmt.Sprintf("organizations/%s/events", organization),
		q,
		&result,
	); e != nil {
		return nil, e
	}

	return result.Rows, nil
}
