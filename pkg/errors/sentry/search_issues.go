package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) SearchIssues(
	organization string,
	query string,
	project string,
	limit int,
	cursor string,
) ([]response.Issue, error) {
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
		fmt.Sprintf("organizations/%s/issues", organization),
		q,
	)

	if e != nil {
		return nil, e
	}

	var result []response.Issue

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
