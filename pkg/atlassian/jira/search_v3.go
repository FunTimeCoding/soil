package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
)

func (c *Client) SearchV3(
	query string,
	a ...any,
) ([]*response.Issue, error) {
	// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	// Do not enrich, otherwise watchedIssueKeys will be recursive.
	return c.searchV3Pages(
		query,
		constant.JiraAllFields,
		constant.JiraChangelogExpand,
	)
}
