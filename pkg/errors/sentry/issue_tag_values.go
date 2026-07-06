package sentry

import (
	"encoding/json"
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

	b, e := c.basic.Get(
		fmt.Sprintf(
			"organizations/%s/issues/%s/tags/%s/values",
			organization,
			identifier,
			tag,
		),
		q,
	)

	if e != nil {
		return nil, e
	}

	var result []response.TagValue

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return result, nil
}
