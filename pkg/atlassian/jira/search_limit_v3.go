package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
)

func (c *Client) SearchLimitV3(
	limit int,
	query string,
	a ...any,
) ([]*response.Issue, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, e := c.searchV3Page(
		limit,
		"",
		query,
		constant.JiraAllFields,
		constant.JiraChangelogExpand,
	)

	if e != nil {
		return nil, e
	}

	return result.Issues, nil
}
