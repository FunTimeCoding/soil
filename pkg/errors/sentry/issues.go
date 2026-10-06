package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
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

	var result []response.Issue

	if e := c.basic.Get(
		fmt.Sprintf("organizations/%s/issues", organization),
		query,
		&result,
	); e != nil {
		return nil, e
	}

	return issue.NewSlice(result), nil
}
