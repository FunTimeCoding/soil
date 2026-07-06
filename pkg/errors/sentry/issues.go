package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	"github.com/funtimecoding/soil/pkg/errors/sentry/validate"
)

func (c *Client) Issues(
	organization string,
	projectIdentifier string,
	period string,
) ([]*issue.Issue, error) {
	validate.Contains(constant.Periods, period)
	query := map[string]string{
		"project": projectIdentifier,
		"query":   constant.UnresolvedFilter,
	}

	if period != "" {
		query["statsPeriod"] = period
	}

	b, e := c.basic.Get(
		fmt.Sprintf("organizations/%s/issues", organization),
		query,
	)

	if e != nil {
		return nil, e
	}

	var result []response.Issue

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return issue.NewSlice(result), nil
}
