package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	"strings"
)

func (c *Client) IssueByShortIdentifier(
	organization string,
	identifier string,
) (*issue.Issue, error) {
	b, e := c.basic.Get(
		fmt.Sprintf(
			"projects/%s/%s/issues",
			organization,
			strings.ToLower(
				identifier[:strings.LastIndex(identifier, "-")],
			),
		),
		map[string]string{"shortIdLookup": "1", "query": identifier},
	)

	if e != nil {
		return nil, e
	}

	var result []response.Issue

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"issue not found: %s: %w",
			identifier,
			constant.ErrorNotFound,
		)
	}

	return issue.New(&result[0]), nil
}
