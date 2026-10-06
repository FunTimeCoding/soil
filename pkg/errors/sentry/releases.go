package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) Releases(
	organization string,
	query string,
	limit int,
) ([]response.Release, error) {
	q := map[string]string{}

	if query != "" {
		q["query"] = query
	}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	var result []response.Release

	if e := c.basic.Get(
		fmt.Sprintf("organizations/%s/releases", organization),
		q,
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
