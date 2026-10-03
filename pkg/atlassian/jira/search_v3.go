package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
)

// Reference: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search
func (c *Client) SearchV3(
	query string,
	a ...any,
) ([]*response.Issue, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	return c.searchV3Pages(
		query,
		constant.JiraAllFields,
		constant.JiraChangelogExpand,
	)
}
