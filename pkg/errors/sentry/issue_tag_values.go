package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) IssueTagValues(
	organization string,
	identifier string,
	tag string,
	limit int,
) ([]response.TagValue, error) {
	q := map[string]string{}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	var result []response.TagValue

	if e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s/tags/%s/values",
			organization,
			identifier,
			tag,
		),
		q,
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
