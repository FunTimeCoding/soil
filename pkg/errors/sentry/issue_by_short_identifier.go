package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	"strings"
)

func (c *Client) IssueByShortIdentifier(
	organization string,
	identifier string,
) (*issue.Issue, error) {
	var result []response.Issue

	if e := c.basic.Get(
		fmt.Sprintf(
			"projects/%s/%s/issues",
			organization,
			strings.ToLower(identifier[:strings.LastIndex(identifier, "-")]),
		),
		map[string]string{"shortIdLookup": "1", "query": identifier},
		&result,
	); e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, not_found.New("issue", identifier)
	}

	return issue.New(&result[0]), nil
}
