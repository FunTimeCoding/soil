package sentry

import (
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

	var result []response.Issue

	if e := c.basic.Get(
		fmt.Sprintf("organizations/%s/issues", organization),
		q,
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
